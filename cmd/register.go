package cmd

import (
	"fmt"
	"HW7/entity"
	"github.com/spf13/cobra"
)

var username string
var password string
var email string
var phone string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "YQC's register",
	Long: `for example: go run main.go register -n user1 -s 1234 -e 123456@qq.com -p 1234567`,
	Run: func(cmd *cobra.Command, args []string) {
		var flagC int = entity.CheckRegister(username,password,email,phone)
		if flagC == 1 {
			fmt.Println("创建成功！")
		} else if flagC == 2 {
			fmt.Println("用户重名！")
		} else if flagC == 3 {
			fmt.Println("未输入用户名！")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&username, "username", "n", "", "username")
	registerCmd.Flags().StringVarP(&password, "password", "s", "123456", "password")
	registerCmd.Flags().StringVarP(&email, "email", "e", "123456@qq.com", "email")
	registerCmd.Flags().StringVarP(&phone, "phone", "p", "123456", "phone")//倒数第二个参数是未填写时的默认值
}
