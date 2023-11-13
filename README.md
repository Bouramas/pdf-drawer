# PDF Drawer

This is a simple Go application that allows you to draw 
either a certificate or an invoice based on command-line 
flags. It is my take on Jonathan Calhoun's [Exercise #20: Building PDFs](https://courses.calhoun.io/courses/cor_gophercises).

## Usage

### Flags

- `-name`: Specifies the name of the person who completed the course.
- `-certificate`: Draws a certificate if this flag is provided.
- `-invoice`: Draws an invoice if this flag is provided.


### Example

To draw a certificate:

```bash
go run main.go -certificate -name JohnDoe
```

To draw an invoice:

```bash
go run main.go -invoice
```