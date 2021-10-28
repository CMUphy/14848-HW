import logging
import sys
import webbrowser

import os
logging.basicConfig(stream=sys.stdout,level=logging.DEBUG)
JUPYTER_NOTEBOOK = os.getenv('JUPYTER_NOTEBOOK', '')

while(True):
    logging.info("Welcome to Big Data Processing Application\n"
                 "Please type the number that corresponds to which application you would like to run:\n"
                 "1. Apache Hadoop\n"
                 "2. Apache Spark\n"
                 "3. Jupyter Notebook\n"
                 "4. SonarQube and SonarScanner\n"
                 )
    try :
        str = input("Typer the number here > \n")
    except Exception as e:
        continue
    if(str == "1"):
        url = JUPYTER_NOTEBOOK + ":81"
        # url1 ="http://www.baidu.com"
        webbrowser.open(url)