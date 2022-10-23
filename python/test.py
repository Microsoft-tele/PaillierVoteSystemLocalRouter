from selenium.webdriver import Chrome
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
import time

web = Chrome()
web.get("http://localhost:8080/index")
#web.get("http://192.168.1.105:8080/index")
time.sleep(1)
element = web.find_element(By.XPATH, '/html/body/div/div[4]/a/button').click()

time.sleep(1)

for i in range(8):
    element = web.find_element(By.XPATH, '/html/body/div/div/form/input[1]').click()

namelist = ["李为君","何俭涛","李林轩","胡靛青","李文豪","沈文涛","熊琦","徐许越","闵浩哲"]

for i in range(8):
    element = web.find_element(By.XPATH, f'//*[@id="Ticket"]/tr[{i+1}]/td[1]/input')
    element.send_keys(namelist[i])
    element = web.find_element(By.XPATH, f'//*[@id="Ticket"]/tr[{i+1}]/td[2]/input')
    element.send_keys(f"{namelist[i]}的自我介绍")


time.sleep(1)

element = web.find_element(By.XPATH, '/html/body/div/div/form/input[2]').click()

#web = web.get("http://localhost:12345/index")

web.close()