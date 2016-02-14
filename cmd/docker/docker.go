package docker
import (
	"github.com/codeskyblue/go-sh"
	"strconv"
	"strings"
	"github.com/tahitianstud/ata/config"
)

var logger = config.Logger

// ActiveContainersCount returns the number of active containers for the given application
func ActiveContainersCount(app string) (int, error) {

	//docker ps -a -q --filter=name=${__FILTER} | wc -l

	out, err := sh.Command("docker", "ps", "-a", "-q", "--filter=name="+app).Command("wc", "-l").Output()

	if err != nil {
		return -1, err
	} else {
		outputString := strings.Replace(string(out), "\n", "", -1)
		result, outerr := strconv.ParseInt(outputString, 10, 64)
		return int(result), outerr
	}
}

// GetContainersList returns the string listing the containers
func GetContainersList(app string, idsOnly bool) (string, error) {

	//docker ps -a -q --filter=name=${__FILTER}

	var (
		out []byte
		err error
	)

	if idsOnly {
		out, err = sh.Command("docker", "ps", "-a", "-q", "--filter=name=" + app).Output()
	} else {
		out, err = sh.Command("docker", "ps", "-a", "--filter=name=" + app).Output()
	}

	if err != nil {
		return "", err
	} else {
		return string(out), err
	}
}
