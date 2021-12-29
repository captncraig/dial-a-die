dial_hole_d=110;
spacer = 15;
display_x = 66.8;
display_y = 88;
tot_x = dial_hole_d + 3*spacer+display_x;
tot_y = dial_hole_d + 2*spacer;
dial_x = dial_hole_d/2+spacer+spacer+display_x;
dial_y = dial_hole_d/2+spacer;

base_y = cos(15)*tot_y;

module top_plate(){
translate([0,0,24])
rotate([15,0,0]){
difference(){
    // top plate
    cube([tot_x,tot_y,2]);
    // cutouts
    translate([dial_x,dial_y,-1])cylinder(d=dial_hole_d,h=5);
    dispx = spacer;
    dispy = (tot_y-display_y)/2;
    translate([dispx,dispy,-1])cube([display_x,display_y,5]);
    blx = dispx-6.8;
    bly = dispy-2.8;
    dx = 72;
    dy= 97;
    translate([blx,bly,-1])cylinder(d=3,h=5);
    translate([blx+dx,bly,-1])cylinder(d=3,h=5);
    translate([blx,bly+dy,-1])cylinder(d=3,h=5);
    translate([blx+dx,bly+dy,-1])cylinder(d=3,h=5);
    
}
    // dummy dial block
    //translate([dial_x,dial_y,-23])cylinder(d=dial_hole_d,h=23);
}
}

// whole top frame
module top(){
difference(){
    union(){
        top_plate();
        // front wall
        translate([0,0,5])cube([tot_x,3,20]);
        translate([0,base_y-3,5])cube([tot_x,3,56.5]);
        // base plate
        color("blue")translate([0,0,1.5])cube([tot_x,base_y,3]);
        
        translate([15,5,5])cylinder(d=pi_standoff_d,h=21.2);
        translate([tot_x-15,5,5])cylinder(d=pi_standoff_d,h=21.2);
        translate([15,base_y-5,5])cylinder(d=pi_standoff_d,h=54);
        translate([tot_x-15,base_y-5,5])cylinder(d=pi_standoff_d,h=54);
    }
    // lop off front
    translate([-1,-5,1])cube([tot_x+10,5,tot_x+10]);
    
    translate([15,5,0])cylinder(d=pi_hole_d,h=20);
    translate([tot_x-15,5,0])cylinder(d=pi_hole_d,h=20);
    translate([15,base_y-5,0])cylinder(d=pi_hole_d,h=20);
    translate([tot_x-15,base_y-5,0])cylinder(d=pi_hole_d,h=20);
}
}
top();

bracket_x = 4;
bracket_y = 25;
bracket_z = 34;
bracket_xpos=dial_x;
bracket_ypos=base_y/2;
mounting_space = 89;
plate_z = 3;
slot_d = 7;

module bracket(){
    linear_extrude(height = bracket_z, scale = [1,.35],center=true) {
        square(size = [bracket_x, bracket_y],center=true);
    }
}
// dial mount brackets
difference(){
    translate([-(mounting_space/2+bracket_x/2)+bracket_xpos,bracket_ypos,bracket_z/2+plate_z])
    {
        bracket();
        translate([mounting_space+bracket_x,0,0])bracket();
    }
    translate([bracket_xpos,bracket_ypos,bracket_z+plate_z-slot_d+.01])rotate([0,90,0])cylinder(d=4.1,h=100,center=true);//cube([mounting_space+3*bracket_x,slot_d,slot_d],center=true);
}

pi_x = 23;
pi_y = 58;
pi_standoff_d = 9;
pi_hole_d = 3;
pi_standoff_h = 10;
module pimount(){
    for (x = [0, pi_x]){for (y = [0,pi_y]){
        translate([x, y, plate_z])difference(){cylinder(d = pi_standoff_d, h = pi_standoff_h);cylinder(d = pi_hole_d, h = pi_standoff_h);}
    }}
    
}
translate([25,35,0])pimount();