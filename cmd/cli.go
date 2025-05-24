package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project-app-todo-list-cli/service"
)

func RunCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== TO-DO LIST MENU ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Tugas")
		fmt.Println("3. Update Status Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu [1-5]: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Masukkan judul tugas: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Masukkan prioritas (low/medium/high): ")
			priority, _ := reader.ReadString('\n')
			err := service.AddTask(strings.TrimSpace(title), strings.TrimSpace(priority))
			if err != nil {
				fmt.Println("Gagal menambahkan tugas")
			} else {
				fmt.Println("Tugas berhasil ditambahkan!")
			}
		case "2":
			service.ListTasks()
		case "3":
			fmt.Print("Masukkan nomer yang akan diupdate:")
			idxstr, _ := reader.ReadString('\n')
			fmt.Println("Masukkan status baru (new/progree/completed): ")
			status, _ := reader.ReadString('\n')

			index, err := strconv.Atoi(strings.TrimSpace(idxstr))
			if err != nil {
				fmt.Println("Input Data Tidak Valid")
				continue
			}

			err = service.UpdateTaskStatus(index, strings.TrimSpace(status))
			if err != nil {
				fmt.Println("Gagal Update:", err)
			} else {
				fmt.Println("Task Updated Succesfully")
			}
		case "4":
			fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
			idxstr, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(idxstr))
			if err != nil {
				fmt.Println("Input tidak valid")
				continue
			}
			err = service.DeleteTask(index)
			if err != nil {
				fmt.Println("Gagal Menghapus:", err)
			} else {
				fmt.Println("Tugas berhasil dihapus.")
			}
		case "5":
			fmt.Println("Terima Kasih, Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
