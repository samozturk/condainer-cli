package environ

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// For error logging
type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;32m%s\033[0m" // Green
	WarningColor = "\033[1;33m%s\033[0m" // Yellow
	ErrorColor   = "\033[1;31m%s\033[0m" // Red
	// Refer to https://www.shellhacks.com/bash-colors/ for different colors.
)

const envBindDir string = "/home/tazi/tmp/envs"

// const contHomeDir string = "/home/tazi"

func ShowMessage(messageType messageType, message string) {
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

func CreateEnv(containerName string, envName string) (string, error) {
	shortCmd := fmt.Sprintf("/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/%v python=3.7.10 pip", envName)
	cmdStr := fmt.Sprintf("docker exec %v /bin/bash -c %q", containerName, shortCmd)
	//cmdStr := fmt.Sprintf("docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/%v python=3.7.10 pip'", containerName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func CloneEnv(containerName string, envName string, newEnvName string) (string, error) {

	command := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y --name %v --clone %v'"
	cmdStr := fmt.Sprintf(command, containerName, newEnvName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func RemoveEnv(containerName string, envName string) (string, error) {

	command := "docker exec %v /bin/bash -c '/home/tazi/miniconda3/bin/conda env remove -n %v'"
	cmdStr := fmt.Sprintf(command, containerName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func AddPackage(containerName string, envName string, packageName string) (string, error) {
	var command string

	command = "docker exec %v bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v'"
	cmdStr := fmt.Sprintf(command, containerName, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func RemovePackage(containerName string, envName string, packageName string) (string, error) {
	command := "docker exec %v bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip uninstall -y %v'"
	cmdStr := fmt.Sprintf(command, containerName, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

func UpdatePackage(containerName string, envName string, packageName string) (string, error) {

	command := "docker exec %v bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install %v --upgrade'"
	cmdStr := fmt.Sprintf(command, containerName, envName, packageName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err
}

// DONE
func AddZipPackage(envName string, source string) error {
	// Get file extension
	fileExt := strings.TrimPrefix(filepath.Ext(source), ".")
	// Get home directory
	homedir, hErr := os.UserHomeDir()
	if hErr != nil {
		log.Fatal(hErr)
	}
	// Fix python3.7, it doesnt have to be like that always.
	dest := fmt.Sprintf("%v/tmp/envs/%v/lib/python3.7/site-packages", homedir, envName)
	if fileExt == "zip" {
		ShowMessage(WARNING, dest)
		UnzipSource(source, dest)
		return nil
	} else if fileExt == "tar" {
		ShowMessage(WARNING, dest)
		Untar(source, dest)
	}
	return nil
}

func AddFromText(containerName string, envName string, source string) (string, error) {
	// Copy requirements.txt to miniconda3/envs/ which is a shared directory
	dst := fmt.Sprintf("%v/requirements.txt", envBindDir)
	CopyFile(source, dst)
	// activate environment name and execute pip install requirements.txt √
	command := "docker exec %v bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate %v; pip install -r /home/tazi/miniconda3/envs/requirements.txt'"
	cmdStr := fmt.Sprintf(command, containerName, envName)
	infoMessage := fmt.Sprintf("Running the command: %v", cmdStr)
	ShowMessage(WARNING, infoMessage)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	sOut := fmt.Sprintf("%s", out)
	return sOut, err

}

func main() {
	// CopyFile("/Users/samet/Documents/Projects/rte-cli/environ/rte-cli", "/Users/samet/Documents/rte-clu")

}

// Util functions
func CopyFile(src string, dst string) {
	// Open original file
	original, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	// Create new file
	new, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
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
