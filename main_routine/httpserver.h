#include "FS.h"
#include <WebServer.h>
// ComponentInterface ensures all components can keep track of their
// initialized states
class ComponentInterface {
protected:
  bool initialized;
};

class FileController : protected ComponentInterface {
public:
  FileController();
  char* readHtml(fs::FS &fs, const char* path);
  bool writeJpg(fs::FS &fs, const char* jpgFilename, const char* fileData);
};

class HttpController : protected ComponentInterface {
public:
  HttpController();

private:
  WebServer serv;
};
