<p align="center">
  <img src="images/index.png" width="400px">
  <img src="images/note.png" width="400px">
</p>
<h1 align="center">Note Server 📝 Host your markdown notes</h1>
<h3 align="center">
  Just a web server that you can use to host
  your markdown notes - written in Go!
</h3>

<p align="center">
    <img alt="GitHub License" src="https://img.shields.io/github/license/ngn13/note-server?style=for-the-badge">
    <img alt="GitHub actions" src="https://img.shields.io/github/actions/workflow/status/ngn13/note-server/publish.yml?style=for-the-badge">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/ngn13/note-server?style=for-the-badge">
</p>

## ✨ Features
- Minimal dark UI (with no javascript)
- Content and path based searching 
- Modification time based sorting
- Image and link support 

## 😋 Setup
### with Docker 
```bash
# bring in your notes, for example:
git clone https://github.com/your/cool-notes.git
mv cool-notes notes

docker run -d -v $PWD/notes:/notes      \
              -p 80:8080                \
              --name my-notes           \
              ghcr.io/ngn13/note-server 
```
Now you can connect to note server on port 80, you can extend this setup with a reverse proxy if you wish.

### without Docker
You can build and install the application from the source using `make`:
```bash
# obtain the latest sources
git clone https://github.com/ngn13/note-server
cd note-server

make
```

To install the binary to your PATH, you can use run the installation command as root:
```bash
sudo make install
```

Then you can run the server with `note-server` command:
```bash
note-server -notes /path/to/your/notes
```

## ⚙️ Options
To list available command-line options, use the `-help` flag:
```
Usage of note-server:
  -interface string
        Web server interface (host:port) (default "127.0.0.1:8080")
  -notes string
        Path for the directory that contains your notes
  -static string
        Static files directory path (default "/usr/lib/note-server/static")
  -views string
        HTML templates directory path (default "/usr/lib/note-server/views")
```

## 🔄 Auto-updating your notes 
If you are using git for your notes, then you can setup a cronjob to auto-update your notes.

To do this add this enrty to your `/etc/crontab`:
```
  0  *  * * *   your-username     cd /path/to/your/notes && /usr/bin/git pull
```
This entry will pull and sync your notes with the remote every hour.

## 🔗 Credit
- [github-markdown-css](https://github.com/sindresorhus/github-markdown-css): Used for markdown rendering
- [nerdfonts](https://github.com/ryanoasis/nerd-fonts): Fonts used in the application 
