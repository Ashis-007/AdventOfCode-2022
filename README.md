### Run program
```bash
cd 02
go run main.go --file=input.txt --part=1
```

### Flags
- file: Input file name. Possible values are `input.txt, demo.txt`
- part: Question part. Possible values are `<1, 2>`

### Template
I have created a shell script `new-day.sh` to generate the template for a new day problem which:
- Initializes a go module based on the arguement passed
- Creates a `main.go` file with a basic template
- Creates two input files i.e. `demo.txt` and `input.txt`

e.g. To create a template for day 5, run the following cmd:
```bash
./new-day.sh 05
```
