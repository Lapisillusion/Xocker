package container

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func CommitContainer(tarName string) {
	mntPath := "/root/merged"
	imageTar := "/root/" + tarName + ".tar"
	fmt.Println("commitContainer imageTar:", imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntPath, ".").CombinedOutput(); err != nil {
		log.Errorf("tar folder %s error %v", mntPath, err)
	}

}
