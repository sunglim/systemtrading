#!/usr/bin/env python
# coding: utf8

import os
import requests
from html.parser import HTMLParser

class KrxHTMLParser(HTMLParser):
    def __init__(self):
        HTMLParser.__init__(self)
        self.isStarted = False
        self.trStarted = False
        self.tdInsideofTrStarted = False

        self.tdCount = 0

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

    def handle_data(self, data):
        if data.strip() == "":
            return

        if self.tdCount == 1:
            print("Encountered some data  :", data)
        if self.tdCount == 2:
            print("Code[" + data + "]")

class KrxCodeGenerator(object):

  def Generate(self):
    URL = 'http://kind.krx.co.kr/corpgeneral/corpList.do?method=download'
    response = requests.get(URL)
    response.encoding = response.apparent_encoding
    #print(response.text)

    parser = KrxHTMLParser()
    parser.feed(response.text)

if __name__ == '__main__':
  object = KrxCodeGenerator()
  object.Generate()
