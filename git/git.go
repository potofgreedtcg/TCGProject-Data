package git

import (
	"fmt"
	"os/exec"
)

func GitCommit() {
	// git add .
	cmdAdd := exec.Command("git", "add", ".")
	if err := cmdAdd.Run(); err != nil {
		fmt.Println("git add failed:", err)
		return
	}

	// git commit -m "your commit message"
	cmdCommit := exec.Command("git", "commit", "-m", "Automated commit from Go")
	if err := cmdCommit.Run(); err != nil {
		fmt.Println("git commit failed:", err)
		return
	}

	// git push
	cmdPush := exec.Command("git", "push")
	if err := cmdPush.Run(); err != nil {
		fmt.Println("git push failed:", err)
		return
	}

	fmt.Println("Successfully committed and pushed.")
}