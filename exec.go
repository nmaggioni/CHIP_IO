package CHIP_IO

import (
	"golang.org/x/sys/unix"
	"log"
	"os/exec"
	"os/user"
	"strings"
)

func shell(command string) (string, error) {
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func checkRootOrGroup() {
	userDetails, err := user.Current()
	if err != nil {
		log.Fatal("Unable to check current user: " + err.Error())
	}

	if userDetails.Uid == "0" || userDetails.Username == "root" {
		return
	}

	currentUser, _ := user.Current()
	groupIDs, _ := currentUser.GroupIds() // TODO: simplify
	for _, id := range groupIDs {
		group, _ := user.LookupGroupId(id)
		if group.Name == "dialout" || group.Name == "gpio" {
			err := unix.Access("/sys/class/gpio", unix.W_OK)
			if err == nil {
				return
			}
		}
	}

	log.Fatal("CHIP_IO needs root privileges or UDEV rule to interact with GPIO! " +
		"See https://github.com/nmaggioni/CHIP_IO/#about-root-privileges")
}
