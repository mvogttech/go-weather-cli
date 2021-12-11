## A Simple Weather CLI

![go-weather-cli in action](https://github.com/mvogttech/go-weather-cli/blob/main/action.png?raw=true "The weather... in your terminal!")

After being tired of going to different weather sources / websites filled with ads, I decided to create a simplistic free-to-use weather forecast report that I could access right in my terminal.

## Where Does The Information Come From?

The weather forecast originates from weather.gov's public free-to-use API. Meaning this software will be free and accessible to anyone.

**That means there is no extra setup or API key configuration. Just install the app and run from your terminal.**

## Usage

- Download the [executable](https://github.com/mvogttech/go-weather-cli/releases/download/V1/go-weather-cli "Latest go-weather-cli release")
- Move to desired folder on your machine `/Users/youruser/standaloneapps/`
- Setup terminal alias
- Open Terminal
- Open .zshrc `sudo nano ~/.zshrc`
- Create alias `alias weather=/Users/youruser/standaloneapps/go-weather-cli`
- Save file `Ctrl+X`
- User `weather` command from your terminal

## License

MIT License

Copyright (c) 2021 Michael Vogt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
