while getopts r:d: flag
do
    case "${flag}" in
        r) input=${OPTARG};;
        d) wheelhouse=${OPTARG};;
    esac
done
​
declare -a FAILED_DEPS=()
green="\033[1;32m"
red="\033[1;31m"
yellow="\033[1;33m"
purple="\033[1;35m"
end="\033[m"
​
PLATFORMS=( "any" "linux_x86_64" "manylinux_2_12_x86_64" "manylinux2010_x86_64" "manylinux2014_x86_64" "manylinux1_x86_64") 
​
while IFS= read -r line
do
  cmd=""
  for plat in ${PLATFORMS[@]}; do
    cmd="$cmd || pip download --only-binary=:all: --platform $plat $line -d $wheelhouse > /dev/null 2>&1"
  done
  cmd=${cmd:4}
  eval "$cmd"
  if [ $? -ne 0 ]
  then
    FAILED_DEPS+=($line)
  fi
done < "$input"
​
(cp $input $wheelhouse/requirements.txt)
(zip $wheelhouse.zip $wheelhouse)
(rm -r $wheelhouse)
​
if [ "${#FAILED_DEPS}" -eq 0 ]; 
  then
    echo "$green Environment zipped successfully. Please upload this zip file to TAZI platform.$end"
  else
    echo "$red Following dependencies couldn't downloaded: $end"
    echo "$purple ${FAILED_DEPS[*]} $end"
    echo "$red Are yo sure their version is available for CentOS environment?$end"
fi
