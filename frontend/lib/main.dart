import 'dart:convert';
import 'dart:io';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'data/model/response.dart';

import 'widget/button.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final TextEditingController inputController = TextEditingController();
  final TextEditingController outputController = TextEditingController();
  String inputText = '';

  @override
  void dispose() {
    // Clean up the controller when the widget is disposed.
    inputController.dispose();
    super.dispose();
  }

  postData(String dataStr) async {
    String result;
    try {
      var response = await Dio(BaseOptions(
        baseUrl: "http://localhost:3000",
        contentType: 'application/json; charset=utf-8',
        connectTimeout: 5000,
        receiveTimeout: 5000,
        responseType: ResponseType.json,
      )).post("/decode", data: {'tx': dataStr});
      if (response.statusCode == HttpStatus.ok) {
        var data = jsonDecode(response.toString());
        var resp = Result.fromJson(data);
        if (resp.code != 0) {
          result = resp.msg!;
        } else {
          result = resp.data != null ? getPrettyJSONString(resp.data) : '';
        }
      } else {
        result = 'Error getting status ${response.statusCode}';
      }
    } catch (exception) {
      result = exception.toString();
    }
    setState(() {
      outputController.text = result;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Decode Ethereum serialized transaction"),
      ),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            padding: const EdgeInsets.all(10),
            margin: const EdgeInsets.only(top: 50.0, right: 200.0),
            child: TextField(
              autofocus: true, // 在文本框可见时将其聚焦
              controller: inputController,
              maxLines: 12,

              decoration: const InputDecoration(
                enabledBorder: OutlineInputBorder(
                  borderSide: BorderSide(width: 1, color: Colors.black54),
                ),
                focusedBorder: OutlineInputBorder(
                  borderSide: BorderSide(width: 1, color: Colors.blueAccent),
                ),
                filled: false,
                fillColor: Color.fromARGB(255, 224, 226, 230),
                hintStyle: TextStyle(color: Color.fromARGB(255, 10, 11, 11)),
                hintText: "Enter serialized ethereum transaction",
              ),
            ),
          ),
          Container(
            padding: const EdgeInsets.fromLTRB(10, 5, 10, 2),
            margin: const EdgeInsets.only(top: 1.0, left: 5, right: 200.0),
            child: Row(
              children: [
                Button("Decode", onPressed: (() {
                  //TODO
                  print("decode pressed");
                  print(inputController.text);
                  postData(inputController.text);
                })),
                Button("Broadcast", onPressed: (() {
                  print("broadcast pressed");
                }))
              ],
            ),
          ),
          Container(
            padding: const EdgeInsets.all(10),
            margin: const EdgeInsets.only(top: 1.0, right: 200.0),
            child: TextField(
              autofocus: true, // 在文本框可见时将其聚焦
              controller: outputController,
              maxLines: 15,

              decoration: const InputDecoration(
                enabledBorder: OutlineInputBorder(
                  borderSide: BorderSide(width: 1, color: Colors.black54),
                ),
                focusedBorder: OutlineInputBorder(
                  borderSide: BorderSide(width: 1, color: Colors.blueAccent),
                ),
                filled: false,
                fillColor: Color.fromARGB(255, 224, 226, 230),
                hintStyle: TextStyle(color: Color.fromARGB(255, 10, 11, 11)),
                //hintText: "Enter serialized ethereum transaction",
              ),
            ),
          ),
        ],
      ),
    );
  }
}
