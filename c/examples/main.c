#include <stdlib.h>     //exit()
#include <signal.h>     //signal()
#include "EPD_4in2.h"
#include "EPD_Test.h"   //Examples

int main(int argc, char *argv[])
{

    if(DEV_Module_Init()!=0){
        return -1;
    }

    char* fname;
    if (argc == 2){
        fname = argv[1];
    }else{
        EPD_4IN2_Init_4Gray();
        return 0;
    }
	
    UBYTE *BlackImage;
    UWORD Imagesize = ((EPD_4IN2_WIDTH % 8 == 0)? (EPD_4IN2_WIDTH / 4 ): (EPD_4IN2_WIDTH / 4 + 1)) * EPD_4IN2_HEIGHT;
    
    if((BlackImage = (UBYTE *)malloc(Imagesize)) == NULL) {
        printf("Failed to apply for black memory...\r\n");
        return -1;
    }
    Paint_NewImage(BlackImage, EPD_4IN2_WIDTH, EPD_4IN2_HEIGHT, 0, WHITE);
	Paint_SetScale(4);
    GUI_ReadBmp_4Gray(fname,0 , 0);
	EPD_4IN2_4GrayDisplay(BlackImage);
    return 0;
}
