// dial bracket
dial_hole_d=110;
mounting_space = 89;
bracket_x = 4;
bracket_y = 25;
bracket_z = 25;
bracket_xpos=110;
bracket_ypos=40;
slot_d = 3.8;

// rpi numbers
pi_x = 23;
pi_y = 58;
pi_standoff_d = 7;
pi_hole_d = 2.5;
pi_standoff_h = 6;

//base plate
plate_x = bracket_xpos+mounting_space/2+15;
plate_y = pi_y + 30;
plate_z = 2;

// base plate
cube([plate_x,plate_y,plate_z]);

// dial mount brackets
difference(){
translate([-(mounting_space/2+bracket_x/2)+bracket_xpos,bracket_ypos,bracket_z/2+plate_z])
{
bracket();
translate([mounting_space+bracket_x,0,0])bracket();
}
translate([bracket_xpos,bracket_ypos,bracket_z+plate_z-slot_d/2+.01])cube([mounting_space+3*bracket_x,slot_d,slot_d],center=true);
}

// pi standoffs
translate([15,15,0])pimount();

//dial standin
module fakedial(){
axle_to_top = 16;
axle_to_plate = 5.5;
dial_plate_d = 114;
dial_angle = 25;
translate([bracket_xpos,bracket_ypos,bracket_z+plate_z-slot_d/2])
rotate([dial_angle,0,0])
{
rotate([0,90,0])cylinder(d=5,h=mounting_space+4*bracket_x,center=true);
cylinder(d=dial_hole_d,h=axle_to_top);
cylinder(d=dial_plate_d,h=axle_to_plate);
}
}

module bracket(){
linear_extrude(height = bracket_z, scale = [1,.35],center=true) {
square(size = [bracket_x, bracket_y],center=true);
}
}

module pimount(){
    for (x = [0, pi_x]){for (y = [0,pi_y]){
        translate([x, y, plate_z])difference(){cylinder(d = pi_standoff_d, h = pi_standoff_h);cylinder(d = pi_hole_d, h = pi_standoff_h);}
    }}
    
}
