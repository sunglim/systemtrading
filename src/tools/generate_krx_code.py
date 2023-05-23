#!/usr/bin/env python

import os
import urllib.request

class KrxCodeGenerator(object):

  def Generate(self):
    response = urllib.request.urlopen('http://kind.krx.co.kr/corpgeneral/corpList.do?method=download')
    html_doc = response.read()
    print(html_doc)

if __name__ == '__main__':
  object = KrxCodeGenerator()
  object.Generate()
