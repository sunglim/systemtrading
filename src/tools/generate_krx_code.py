#!/usr/bin/env python
# coding: utf8

import os
import requests
from html.parser import HTMLParser

class Code(object):
  """A convenience object for constructing code.

  Logically each object should be a block of code. All methods except |Render|
  and |IsEmpty| return self.
  """
  def __init__(self, indent_size=2, comment_length=80):
    self._code = []

  def Append(self, code):
    self._code.append(code)
    return self

  def Render(self):
    """Renders Code as a string.
    """
    return '\n'.join(self._code)

class KrxHTMLParser(HTMLParser):
    def __init__(self):
        HTMLParser.__init__(self)
        self.isStarted = False
        self.trStarted = False
        self.tdInsideofTrStarted = False

        self.tdCount = 0

        self.listOfCode = []
        self.name = ""
        self.code = ""

    def handle_starttag(self, tag, attrs):
        if self.trStarted == True and tag == 'td':
           self.tdInsideofTrStarted = True
           self.tdCount = self.tdCount + 1

        if tag == 'tr':
            self.trStarted = True

    def handle_endtag(self, tag):
        if tag == 'tr':
            self.trStarted = False
            self.tdInsideofTrStarted = False
            self.tdCount = 0
            self.listOfCode.append([self.name, self.code])
            self.name = ""
            self.code = ""

    def handle_data(self, data):
        if data.strip() == "":
            return

        if self.tdCount == 1:
            self.name = data.replace("&", "n").replace("-", "_").replace(" ", "_").replace(".", "")
           # print("name :", data)
        if self.tdCount == 2:
            self.code = data
           # print("Code[" + data + "]")

class KrxCodeGenerator(object):


  def Generate(self):
    URL = 'http://kind.krx.co.kr/corpgeneral/corpList.do?method=download'
    response = requests.get(URL)
    response.encoding = response.apparent_encoding
    #print(response.text)

    parser = KrxHTMLParser()
    parser.feed(response.text)
    #print(parser.listOfCode)

    c = Code()
    c.Append("package koreaexchange")
    c.Append("")
    c.Append("const (")
    for x in parser.listOfCode:
      if x[1] == "":
         continue
      c.Append("  Code" + x[0] + " = \"" + x[1] +"\"")
    c.Append(")")

    print(c.Render())

if __name__ == '__main__':
  object = KrxCodeGenerator()
  object.Generate()
