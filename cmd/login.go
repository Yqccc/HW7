package cmd

import (
	"fmt"
	"HW7/entity"
	"github.com/spf13/cobra"
)
var Username string
var Password string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "YQC's login",
	Long: `for example: go run main.go login -N user1 -S 1234`,
	Run: func(cmd *cobra.Command, args []string) {
		var flagD int = entity.CheckLogin(Username,Password)
		if flagD == 1 {
			fmt.Println("登录成功！")
		} else if flagD == 2 {
			fmt.Println("用户名不存在！")
		} else if flagD == 3 {
			fmt.Println("密码错误！")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&Username, "Username", "n", "", "username")
	loginCmd.Flags().StringVarP(&Password, "Password", "s", "", "password")//倒数第二个参数是未填写时的默认值
}
