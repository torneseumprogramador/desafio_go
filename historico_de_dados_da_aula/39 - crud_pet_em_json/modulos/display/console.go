package display

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func LerDados() string {
	reader := bufio.NewReader(os.Stdin)
	nomeComposto, _ := reader.ReadString('\n')
	return strings.TrimSpace(nomeComposto)
}

func LerDadosComMensagem(mensagem string) string {
	fmt.Println(mensagem)
	return LerDados()
}

func Espera(segundos int) {
	time.Sleep(time.Duration(segundos) * time.Second)
}

func LimparTela() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
