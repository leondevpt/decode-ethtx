// 自定义button按钮
import 'package:flutter/material.dart';

class Button extends StatelessWidget {
  final String text; // 按钮文字
  void Function()? onPressed; // 点击按钮触发的方法
  Button(this.text, {Key? key, required this.onPressed}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      style: ButtonStyle(
        backgroundColor:
            MaterialStateProperty.all(const Color.fromARGB(242, 255, 244, 244)),
        foregroundColor: MaterialStateProperty.all(Colors.black87),
      ),
      onPressed: onPressed,
      child: Text(text),
    );
  }
}
