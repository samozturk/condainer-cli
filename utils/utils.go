package utils

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// using `which conda` to get conda homedir??

type MessageType int

const (
	INFO MessageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;32m%s\033[0m" // Green
	WarningColor = "\033[1;33m%s\033[0m" // Yellow
	ErrorColor   = "\033[1;31m%s\033[0m" // Red
	// Refer to https://www.shellhacks.com/bash-colors/ for different colors.
)

func ShowMessage(messageType MessageType, message string) {
	switch messageType {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		log.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		log.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		log.Printf(ErrorColor, printMessage)
	}
}

// Runs given command and returning stdout, stderr and err
func RunCommand(mainCmd string) (string, string, error) {
	cmd := exec.Command("/bin/bash", "-c", mainCmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil && stderr.String() != "" {
		infoMessage := fmt.Sprint(err) + ": " + stderr.String()
		ShowMessage(ERROR, infoMessage)
		return out.String(), stderr.String(), err
	}
	fmt.Println("Result: " + out.String())
	return out.String(), stderr.String(), err
}

// Util functions
// Parse python versions to use in directories
func GetPyVersion(version string) (int, int, int) {
	m, _ := regexp.Compile("[0-9]+")
	matches := (m.FindAllStringSubmatch(version, -1))
	if len(matches) != 3 {
		ShowMessage(ERROR, "Semantic versioning should be in 'major.minor.patch' format.")
		log.Fatal("Wrong Python version formatting.")
	}
	major, err := strconv.Atoi(matches[0][0])
	if err != nil {
		ShowMessage(ERROR, "Major version can't be parsed.")
	}
	minor, err := strconv.Atoi(matches[1][0])
	if err != nil {
		ShowMessage(ERROR, "Minor version can't be parsed.")
	}
	patch, err := strconv.Atoi(matches[2][0])
	if err != nil {
		ShowMessage(ERROR, "Patch version can't be parsed.")
	}
	return major, minor, patch
}

func CopyFile(src string, dst string) {
	// Open original file
	original, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened source")
	defer original.Close()

	// Create new file
	new, err := os.Create(os.ExpandEnv(dst))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created destination")
	defer new.Close()

	//This will copy
	bytesWritten, err := io.Copy(new, original)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", bytesWritten)
}

func UnzipSource(source, destination string) error {
	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		log.Println(f.Name)
		err := unzipFile(f, destination)
		if err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		ShowMessage(ERROR, "invalid file path")
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create a directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}

	return nil
}

func Untar(sourcefile, dest string) {
	// Check if file path is not empty.
	if sourcefile == "" {
		ShowMessage(ERROR, "Can't find tar file.")
		os.Exit(1)
	}
	// Open the file
	file, err := os.Open(sourcefile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			ShowMessage(ERROR, err.Error())
			os.Exit(1)
		}
		defer fileReader.Close()
	}
	// Define a tar reader
	tarBallReader := tar.NewReader(fileReader)

	// Extract tarred files
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		// get the individual filename and extract to the current directory
		filename := filepath.Join(dest, header.Name)

		// Regarding to the type of indivual file, behave differently
		switch header.Typeflag {
		// check here for constants: https://pkg.go.dev/archive/tar#pkg-constants
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		case tar.TypeReg:
			// handle normal file
			untarMsg := fmt.Sprintf("Untarring : %v", filename)
			ShowMessage(INFO, untarMsg)

			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

// MOVE THIS TO PKGS
// Getting packages for offline use from HOST
func SaveLocalPackages(envName string, dest string) {
	_, _, err := GetReqFileFromLocalEnv(envName, dest)
	if err != nil {
		os.Exit(1)
	}
	cmdStr := fmt.Sprintf("mkdir %v/wheelhouse ; pip download -r %v/requirements.txt -d %v/wheelhouse", dest, dest, dest)
	RunCommand(cmdStr)
	ZipWriter("pymodules/wheelhose", "pymodules/wheelhouse.zip")
}

func GetReqFileFromLocalEnv(envName string, dest string) (string, string, error) {
	cmdStr := fmt.Sprintf("mkdir pymodules; source activate %v; pip freeze > %v/requirements.txt", envName, dest)
	out, stderr, err := RunCommand(cmdStr)
	return out, stderr, err
}

// For zipping files
func ZipWriter(baseFolder string, dest string) {
	// Add slash just in case
	baseFolder = baseFolder + "/"
	// Get a Buffer to Write To
	outFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// We use this to iterate on files recursively to generate folders too:
func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}
	// Iterate through files
	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			// Read binary data of the file
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Create some files in the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			// Write the binary data to created files in the archive
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
