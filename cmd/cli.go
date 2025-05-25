package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project-app-todo-list-cli/service"
	"project-app-todo-list-cli/utils"
)

func RunCLI() {
	//membuat flag
	add := flag.String("add", "", "Tambahkan tugas baru, contoh: --add=\"Belajar Pemograman Golang\"")
	list := flag.Bool("list", false, "Tampilkan semua tugas")
	done := flag.Int("done", 0, "Tandai tugas sebagai selesai, contoh: --done=2")
	del := flag.Int("delete", 0, "Hapus Tugas berdasarkan index, contoh: --delete=2")
	priority := flag.String("priority", "low", "Prioritas tugas: low,medium,high")
	search := flag.String("search", "", "Cari tugas berdasarkan kata kunci")

	flag.Parse()

	//jika ada flag yg digunakan, jlnkan mode cli
	if flag.NFlag() > 0 {
		executed := false
		if *add != "" {
			err := service.AddTask(*add, *priority)
			if err != nil {
				fmt.Println("Gagal menambahkan:", err)
			} else {
				fmt.Println("Tugas berhasil ditambahkan.")
			}
			executed = true
		}

		if *list {
			err := service.ListTasks()
			if err != nil {
				fmt.Println("Gagal menampilkan tugas:", err)
			}
			executed = true
		}

		if *done > 0 {
			err := service.MarkTaskDone(*done)
			if err != nil {
				fmt.Println("Gagal menyelesaikan tugas:", err)
			} else {
				fmt.Println("Tugas ditandai sebagai selesai.")
			}
			executed = true
		}

		if *del > 0 {
			err := service.DeleteTask(*del)
			if err != nil {
				fmt.Println("Gagal menghapus tugas:", err)
			} else {
				fmt.Println("Tugas berhasil dihapus.")
			}
			executed = true
		}

		if *search != "" {
			err := service.SearchTask(*search)
			if err != nil {
				fmt.Println("Gagal mencari tugas:", err)
			}
			executed = true
		}

		if !executed {
			fmt.Println("Perintah tidak valid. Gunakan salah satu opsi:")
			flag.PrintDefaults()
		}

		return
	}

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
			priority = strings.TrimSpace(priority)

			if !utils.IsValidPriority(priority) {
				fmt.Println("Prioritas tidak valid. Gunakan: low, medium, atau high.")
				continue
			}

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
			fmt.Print("Masukkan status baru (new/progress/completed): ")
			status, _ := reader.ReadString('\n')
			status = strings.TrimSpace(status)

			index, err := strconv.Atoi(strings.TrimSpace(idxstr))
			if err != nil {
				fmt.Println("Input Data Tidak Valid")
				continue
			}

			//validasi untuk update status
			/*
				validStatuses := []string{"new", "progress", "completed"}
				isValid := false
				for _, v := range validStatuses {
					if status == v {
						isValid = true
						break
					}
				}
				if !isValid {
					fmt.Println("Status tidak valid.Gunakan salah satu: new,progress,completed")
					continue
				}
			*/
			statusClean := strings.TrimSpace(status)
			if !utils.IsValidStatus(statusClean) {
				fmt.Println("Status tidak valid.Gunakan salah satu: new,progress,completed")
				continue
			}

			//jika valid,lakukan update
			err = service.UpdateTaskStatus(index, status)
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
