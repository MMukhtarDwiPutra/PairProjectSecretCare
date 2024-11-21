package cli

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"
)

func (c *cli) UpdateMyAccount() {
	user, _ := utils.GetUserFromContext(c.ctx)

	var username, password, fullName *string

	userInput := helpers.InputAndHandlingText("Masukan username baru (atau tekan Enter untuk melewati): ")
	if userInput != "" {
		username = &userInput
	}

	passwordInput := helpers.InputAndHandlingText("Masukan password baru (atau tekan Enter untuk melewati): ")
	if passwordInput != "" {
		password = &passwordInput
	}

	fullNameInput := helpers.InputAndHandlingText("Masukan nama lengkap baru (atau tekan Enter untuk melewati): ")
	if fullNameInput != "" {
		fullName = &fullNameInput
	}

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	err := c.handler.User.UpdateMyAccount(user.ID, username, password, fullName)

	done <- true
	fmt.Print("\r                \r")

	newUpdatedUser := &entity.Users{ID: user.ID, TokoID: user.TokoID}
	if username != nil {
		newUpdatedUser.Username = *username
	}
	if fullName != nil {
		newUpdatedUser.FullName = *fullName
	}

	c.ctx = utils.SetUserInContext(c.ctx, newUpdatedUser)
	if err != nil {
		fmt.Printf("Gagal mengubah data akun: %v\n", err)
		return
	}

	updatedUser, ok := utils.GetUserFromContext(c.ctx)
	if !ok {
		fmt.Println("Tidak dapat mengambil data akun yang diperbarui.")
		return
	}

	fmt.Println("Data akun berhasil diubah.")
	fmt.Printf("Informasi akun terbaru:\nUsername: %s\nNama Lengkap: %s\n", updatedUser.Username, updatedUser.FullName)
}

func (c *cli) DeleteMyAccount() {
	user, ok := utils.GetUserFromContext(c.ctx)
	if !ok {
		fmt.Print("user not found in context")
	}

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	c.handler.User.DeleteMyAccount(user.ID)

	done <- true
	fmt.Print("\r                \r")

	c.ctx = context.Background()
}
