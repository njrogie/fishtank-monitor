// NOTE: MOVE THE HTML IN ./html TO SD CARD
#include "httpserver.h"

#define FISHTANK_WATER 2
#define TEST_CHEM 3
#define EXTRACT_WASTE 4

#define MOTOR_ON HIGH
#define MOTOR_OFF LOW

void timed_motor(int pin, int delay_msec)
{
  digitalWrite(pin, MOTOR_ON);
  delay(delay_msec);
  digitalWrite(pin, MOTOR_OFF);
}

void setup() {
  pinMode(FISHTANK_WATER, OUTPUT);
  pinMode(TEST_CHEM, OUTPUT);
  pinMode(EXTRACT_WASTE, OUTPUT);

  // use Serial1 to communicate with ESP01, the wifi endpoint
  Serial1.begin(9600); 
}
void loop()
{
  
}
