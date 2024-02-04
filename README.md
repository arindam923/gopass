# GoPass - Secure Password Generator and Manager

GoPass is a simple yet powerful command-line tool written in Golang for generating and managing secure passwords for different websites. It provides the ability to store passwords, generate random ones on demand, and copy them to the clipboard.

## Features

- **Random Password Generation:** Generate strong and random passwords for various websites.
- **Password Storage:** Store passwords securely with timestamps for easy tracking.
- **Clipboard Integration:** Copy passwords to the clipboard with a single command.


## generating a password
`
    ./cli_app generate -l
`
### gives a website name and it will save and generate a password for that website

### show all the password saved
`
    ./cli_app View
`

## Contributing
Contributions are welcome! If you find any bugs or have new features in mind, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- github.com/atotto/clipboard - Clipboard support in Go.
- github.com/spf13/cobra - A Commander for modern Go CLI interactions.
