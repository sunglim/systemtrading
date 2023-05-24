#!/usr/bin/env python

import os
import requests
from html.parser import HTMLParser

class KrxHTMLParser(HTMLParser):
    def handle_starttag(self, tag, attrs):
        print("Encountered a start tag:", tag)

    def handle_endtag(self, tag):
        print("Encountered an end tag :", tag)

    def handle_data(self, data):
        print("Encountered some data  :", data)

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
