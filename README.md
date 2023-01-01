**Description**<br>
Auto saving your code to the server using ssh

<hr>

**Install**<br>
git clone github.com/chapa-ai/ssh

<hr>

**Using**<br>

Step 1:<br>
    In the main.go file, specify the access data to your server - server ip address, password and login

```bash
	a := Constructor(agent.Options{
		Ip:       "addressIp",
		Password: "password",
		Login:    "login",
	})
```

<hr>

Step 2:<br>
    Also in this file write the path to your project that you want to save to this server.

```bash
    Watcher: Watcher{Path: "/Users/anton/Desktop/modem", ExcludeMatch: []string{".idea", ".git"}},
```

<hr>

Step 3:<br>
    Go to the project you want to save and just change something and save those changes(sometimes you have to do it twice) 

<hr>

Step 4:<br>
    The updates you made to your project code should be implemented. See them on the server
<hr>