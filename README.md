# Intro

This repo produces a daily digest about the Go Programming language. It uses [Matcha](https://github.com/piqoni/matcha) rss generator and Github Pages to host it.  

This is here mainly to be used as a template so you can use it for your own feeds. It also supports password protected pages via [staticrypt](https://github.com/robinmoisson/staticrypt)

# How to make your own 

```
git clone --depth 1 git@github.com:piqoni/go-digest my-personal-digest
cd my-personal-digest
rm matcha.db.enc docs/index.html
rm -rf .git
```

Go to https://github.com/new and create your my-personal-digest repository, and follow instructions to push the my-personal-digest directory as a new repository. (git init, commit, push)

Go to the my-personal-digest repository Settings -> Pages -> Branch: main -> /docs 

Go to the my-personal-digest repository Settings -> Secrets and Variables -> Actions
	- Mandatory: New Repository secret called MATCHA_DB_KEY which is used to encrypt the matcha database
	- Mandatory: Click Variables -> New Repository Variable -> MATCHA_CONFIG_YAML with value of a basic matcha configuration

	```
	markdown_dir_path: pre-docs
	markdown_db_path: matcha.db
	feeds:
	  - http://hnrss.org/best
	  - https://feeds.bbci.co.uk/news/rss.xml
	```

	- Optional if you want page to be password protected: Create a new Repository Secret called STATICRYPT_PASSWORD which will be the password that you will use to access your feed. 


