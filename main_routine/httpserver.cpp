#include "httpserver.h"
#define CHECK_INIT(fail_directive) if(!initialized) {fail_directive;}
#define PRINT(msg) if(Serial) Serial.print(msg);
#define PRINTLN(msg) if(Serial) Serial.println(msg);

// ------------------------- FILE CONTROLLER --------------------
#include "SD.h"

FileController::FileController() 
{
  if(!SD.begin()) { 
    Serial.println("Card mount failed");
    initialized = false;
  }
  if(SD.cardType() == CARD_NONE) {
    initialized = false;
  }
  initialized = true;
}

char* FileController::readHtml(fs::FS &fs, const char* path)
{
  CHECK_INIT(return "SD Card failed to initialize")
  // Read Html file into contents, return length of file
  File file = fs.open(path);
  if(!file) return "Null";

  char contents[file.size()];
  for(int i = 0; i < file.size(); i++)
  {
    contents[i] = file.read();
  }
  return contents;
}

bool FileController::writeJpg(fs::FS &fs, const char* jpgFilename, const char* fileData)
{
  CHECK_INIT(return false)
  if(!fs.exists("/images")) fs.mkdir("/images");
  String filePath = String("/images/") + String(jpgFilename);
  File file = fs.open(filePath.c_str(), FILE_WRITE);
  if(!file) return false;

  if(!file.print(fileData)) return false;
  file.close();
  return true;
}


// -------------------------- WEB CONTROLLER --------------------
#include <WiFi.h>
#include <WebServer.h>

HttpController::HttpController() : serv(80)
{
  const char* ssid = "ATTN7DIea2";
  const char* password = "ksb=+2j+z3ib";

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while(WiFi.status() != WL_CONNECTED) {
    PRINT(".")
    delay(1000);
  }
  initialized = true;
}
