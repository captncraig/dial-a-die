#include <Arduino.h>

#define LED 13
#define DIALPULSE 14
#define DIALLATCH 32
#define GNDA 15
#define GNDB 33

void setup() {
  Serial.begin(9600);
  pinMode(LED, OUTPUT);
  pinMode(DIALPULSE, INPUT_PULLUP);
  pinMode(DIALLATCH, INPUT_PULLUP);
  pinMode(GNDA, OUTPUT);
  pinMode(GNDB, OUTPUT);
  digitalWrite(GNDA, LOW);
  digitalWrite(GNDB, LOW);
}

int latchState;
int lastLatch;
ulong lastLatchBounce;

int pulseState;
int lastPulse;
ulong lastPulseBounce;

const ulong bounceDelay = 18;

int reading;
#define DEBOUNCE(pin,state,last,lastBounce) \
  reading=digitalRead(pin); \
  if (reading != last){lastBounce=now;} \
  if ((now - lastBounce) > bounceDelay && reading != state){\
  state = reading;}last=reading;

bool dialing = false;
bool pulseOn = false;
int pulseCount = 0;
void onDial(int);

void loop() {
  ulong now = millis();
  DEBOUNCE(DIALLATCH,latchState,lastLatch,lastLatchBounce)
  DEBOUNCE(DIALPULSE,pulseState,lastPulse,lastPulseBounce)
  
  if(dialing){
    if(latchState){
      dialing = false;
      onDial(pulseCount);
      return;
    }
    if (!pulseOn && pulseState){
      pulseOn = true;
      pulseCount++;
    }
    if (pulseOn && !pulseState){
      pulseOn = false;
    }
  }else{
    if (!latchState){
      pulseCount = 0;
      pulseOn = false;
      dialing = true;
    }
  }
}

void onDial(int n){
  Serial.println(n);
}