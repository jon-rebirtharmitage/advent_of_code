import cv2
import hashlib
import time
import random

from paho.mqtt import client as mqtt_client

broker = '192.168.1.90'
port = 1883
topic = "cronintech/test"
# Generate a Client ID with the publish prefix.
client_id = f'publish-{random.randint(0, 1000)}'
username = 'cronintech'
password = 'cl0udStrif3'


def connect_mqtt():
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
            print("Connected to MQTT Broker!")
        else:
            print("Failed to connect, return code %d\n", rc)

    client = mqtt_client.Client(client_id)
    client.username_pw_set(username, password)
    client.on_connect = on_connect
    client.connect(broker, port)
    return client

def publish(client):
    while True:
        time.sleep(10)
        ran = captureRandom()
        for k,v in ran.items():
            topic = "cronintech/"+ k
            msg = v
            result = client.publish(topic, msg)
            # result: [0, 1]
            status = result[0]
            if status == 0:
                print(f"Send `{msg}` to topic `{topic}`")
            else:
                print(f"Failed to send message to topic {topic}")

def run():
    client = connect_mqtt()
    client.loop_start()
    publish(client)


def captureRandom():
    r = {}
    cap = cv2.VideoCapture(0, cv2.CAP_DSHOW) 

    cap.set(cv2.CAP_PROP_FRAME_WIDTH, 1000)
    cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 800)
    # reading the input using the camera
    result, image = cap.read()
    grayImage = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    ret, bw_img = cv2.threshold(grayImage, 50, 255, cv2.THRESH_BINARY)

    # Cropping an image
    cropped_image = bw_img
    cmb = ""
    cms = 0
    i = 0
    j = 0
    p = 0
    while i < 720:
        while j < 720:
            if bw_img[i, j] == 255 and p != 255:
                #cmb.append("1")
                p = 255
                cms += 1
            elif bw_img[i, j] == 0 and p != 0:
                #cmb.append("0")
                p = 0
            j+=1
        i += 1
        j = 0
    print(cms)
    talktouniverse = ""
    cmb = ""
    cmba = ""
    i = 0
    j = 0
    while i < 64:
        while j < 64:
            if bw_img[i, j] == 255:
                cmb += "1"
            elif bw_img[i, j] == 0:
                cmb += "0"
            if len(cmb) == 8:
                if cmb != "00000000" and cmb != "11111111":
                    talktouniverse += chr(int(cmb,2))
                    cmba += cmb
                cmb = ""
            j+=1
        i += 1
        j = 0
    #print(talktouniverse)
    x = int(cmba, 2)
    result = hashlib.md5(talktouniverse.encode())
    z = str(x)
    #print("Random 2 Digit numbers : ", z[len(z)-2:])
    ra = round(int(z[len(z)-5:])/100)
    r["random_100"] = str(ra)
    #print("Random 3 Digit numbers : ", z[len(z)-3:])
    rb = round(int(z[len(z)-7:])/1000)
    r["random_1000"] = str(rb)
    r["random_hash"] = result.hexdigest()
    y = format(x, 'x')
    print(y)

    print(result.hexdigest())
    return r


 
    # Display cropped image
    #cv2.imshow("cropped", cropped_image)

    cv2.imshow("Test", image)
    # cv2.imshow("Binary", bw_img)
    # cv2.waitKey(0)
    # cv2.destroyAllWindows()
if __name__ == '__main__':
    run()