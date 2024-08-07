package git_commit_msg

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/quarkcloudio/quark-go/v2/pkg/utils/file"
)

var (
	huskyDir            = ".husky/"
	gitCommitMsgFile    = "commit-msg"
	gitCommitMsgContent = `#!/bin/sh
	
	commit_message=$(cat "$1")
	
	# 定义提交信息的格式要求
	pattern="^(feat|fix|docs|style|refactor|perf|test|build|chore|revert|wip|hotfix): .+"
	
	# 检查提交信息是否符合格式
	if ! echo "$commit_message" | grep -qE "$pattern"; then
		echo "无效的commit提交格式"
		echo "请遵循正确的格式提交. 例如: hotfix: 紧急修复XXXX"
		echo "feat: 新功能"
		echo "fix: 修复 bug"
		echo "docs: 仅文档更改"
		echo "style: 不影响代码含义的更改"
		echo "refactor: 既不修复错误也不添加新功能的代码重构"
		echo "perf: 提高性能的代码更改"
		echo "test: 添加缺失的测试或更正现有测试"
		echo "build: 影响构建系统或外部依赖的更改"
		echo "chore: 改变构建流程"
		echo "revert: 恢复先前的提交"
		echo "wip: 正在进行中的工作，用于临时提交"
		echo "hotfix: 用于紧急修复"
		exit 1
	fi`
)

func init() {

	if !isGitInstalled() {
		return
	}

	if !file.IsExist(huskyDir + gitCommitMsgFile) {
		os.WriteFile(huskyDir+gitCommitMsgFile, []byte(gitCommitMsgContent), 0755)
	}

	if !isGitConfigSet() {
		exec.Command("git", "config", "core.hooksPath", huskyDir).Run()
	}

	return
}

func isGitInstalled() bool {

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "git", "--version")
	default: // Linux and macOS
		cmd = exec.Command("sh", "-c", "git --version")
	}

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func isGitConfigSet() bool {

	cmd := exec.Command("git", "config", "--get", "core.hooksPath")
	if output, err := cmd.Output(); err == nil {
		return string(output) == huskyDir
	}

	return false
}
