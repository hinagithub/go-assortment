// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
//     // args()
//     // flagSample()
//     // boolCount()
//     // placeholder()
//     altername()
// }



// // 引数
// func args(){
//     app := &cli.App{
//         Action: func(cCtx *cli.Context) error {
//             fmt.Printf("Hello %q", cCtx.Args().Get(0))
//             return nil
//         },
//     }

//     // greet hello
//     // hello 

//     if err := app.Run(os.Args); err != nil {
//         log.Fatal(err)
//     }
// }

// func flagSample(){
//     var language string

//     app := &cli.App{
//         Flags: []cli.Flag{
//             &cli.StringFlag{
//                 Name:  "lang",
//                 Value: "english",
//                 Usage: "language for the greeting",
//                 Destination: &language,
//             },
//         },
//         Action: func(cCtx *cli.Context) error {
//             name := "someone"
//             if cCtx.NArg() > 0 {
//                 name = cCtx.Args().Get(0)
//             }
//             if cCtx.String("lang") == "jp" {
//                 fmt.Println("こんにちは", name)
//             } else {
//                 fmt.Println("Hello", name)
//             }
//             return nil
//         },
//     }

//     // greet --lang=jp taro
//     // こんにちは taro

//     if err := app.Run(os.Args); err != nil {
//         log.Fatal(err)
//     }
// }

// func boolCount(){
//     var count int

//     app := &cli.App{
//         Flags: []cli.Flag{
//             &cli.BoolFlag{
//                 Name:        "foo",
//                 Usage:       "foo greeting",
//                 Count: &count,
//             },
//         },
//         Action: func(cCtx *cli.Context) error {
//             fmt.Println("count", count)
//             return nil
//         },
//     }

//     // greet --foo --foo --foo
//     // count 3

//     if err := app.Run(os.Args); err != nil {
//         log.Fatal(err)
//     }
// }

// func placeholder (){
//     app := &cli.App{
//         Flags: []cli.Flag {
//            &cli.StringFlag{
//                 Name: "config",
//                 Aliases: []string{"c"},
//                 Usage: "Load configuration from 'FILE'",
//             },
//         },
//     }

//     if err := app.Run(os.Args); err !=nil{
//         log.Fatal(err)
//     }  
// }

// func altername(){
//     app := &cli.App{
//         Flags:[]cli.Flag{
//             &cli.StringFlag{
//                 Name:"lang",
//                 Aliases:[]string{"l"},
//                 Value : "english",
//                 Usage : "language for greeting",
//             },
//         },
//         Action: func(cCtx *cli.Context) error {
//             lang := cCtx.String("lang")
//             if lang == "jp"{
//                 fmt.Println("こんにちは")
//                 return nil
//             }
//             if lang=="en"{
//                 fmt.Println("hello")
//                 return nil
//             }
//             fmt.Println("Available options are ev | jp")
//             return nil
//         },
//     }
//     // cli >>>greet -l jp
//     // こんにちは
//     // cli >>>greet -lang en
//     // hello
//     // cli >>>greet -l fc
//     // Available options are ev | jp
    
//     if err := app.Run(os.Args); err !=nil{
//         log.Fatal(err)
//     }
// }

