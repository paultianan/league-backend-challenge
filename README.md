# League Backend Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ```
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ```
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ```
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ```
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages

## Developer's Note
In this challenge, I created different routes for the different use-cases. Please refer to the following routes.

- `/echo` : For the Echo (given) scenario.
- `/invert` : For the Invert scenario.
- `/flatten` : For the Flatten scenario.
- `/sum` : For the Sum scenario.
- `/multiply` : For the Multiply scenario.

To run this application, you must have a go language installed on your machine.

Once you have your go language installed, go to the root folder of this application and run the go command `go run .` or `go run main.go`

The default port of this app is `8080`, if this port is not available to you, you can change the port on the `main.go`.

Assuming you run the app with no error, the next step is to use curl below. Note that change the `file-path` with the right location of the csv file.

**ECHO**
```
curl -s -F 'file=@/file-path/matrix.csv' "localhost:8080/echo"
```

**INVERT**
```
curl -s -F 'file=@/file-path/matrix.csv' "localhost:8080/invert"
```

**FLATTEN**
```
curl -s -F 'file=@/file-path/matrix.csv' "localhost:8080/flatten"
```

**SUM**
```
curl -s -F 'file=@/file-path/matrix.csv' "localhost:8080/sum"
```

**MULTIPLY**
```
curl -s -F 'file=@/file-path/matrix.csv' "localhost:8080/multiply"
```