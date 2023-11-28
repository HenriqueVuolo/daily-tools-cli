<h1 align="center">Daily Tools CLI</h1>

<br>
<p align="center">
<img src="http://img.shields.io/static/v1?label=STATUS&message=IN%20PROGRESS&color=GREEN&style=for-the-badge"/>
</p>

## :book: About

Daily Tools CLI is a command-line interface (CLI) developed in Go. The purpose of this application is to provide useful tools for daily tasks directly in your computer's terminal.

The motivation behind creating this project was to learn Go dynamically, and so far, it has been successful. Each new feature brings the opportunity to understand new concepts.

Initially, three commands (`weather`, `dice`, and `password`) were implemented, but the plan is to continue expanding the variety of tools, maintaining a constant learning experience.

## :rocket: Techs

- [Golang](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)

## 🛠️ Commands

Weather 🌥️
- Get the current weather information for a specific location

```bash
> daily-tools-cli weather -c London

```


Dice 🎲
- The dice command allows you to simulate the roll of one or more dice. By default, it rolls a six-sided die (d6), but you can customize the number of dice (A), faces (X), and apply a modifier (B).

```bash
> daily-tools-cli dice -A 2 -X 8 B 3

2d8 + 3 = 13
```


Password 🔒
- Generate a random password

```bash
> daily-tools-cli password -l 10 -u -n

c3BkF3jafA
```

<br><br>
<p align="center">
  <a href="https://www.linkedin.com/in/henrique-vuolo-santana">
  <img src="https://img.shields.io/badge/LinkedIn-Henrique%20Vuolo-blue?logo=linkedin"/></a>
</p>
