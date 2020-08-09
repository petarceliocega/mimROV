<!-- Output copied to clipboard! -->

<!-----
NEW: Check the "Suppress top comment" option to remove this info from the output.

Conversion time: 0.616 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β29
* Sun Aug 09 2020 03:18:52 GMT-0700 (PDT)
* Source doc: ROV control over TCP using Go programming language
----->


<h2>ROV (_Remote Operated Vehicle_) control using TCP protocol and Go programming language</h2>


<p>This is small part of a much larger project where Power Line Communication (PLC) was used to send HD video feed and control a submarine ROV from the ground station.

<h2>Introduction</h2>


_Power Line Communication_ is a technology for digital signal transmission through low-voltage power line(s). Water is a medium that poorly propagates EM waves. To control ROV that can dive reasonably deep underwater, you will need to have it connected by cable(s). Ideally only one cable. That’s where PLC comes in as both the current needed to power the motors and electronics, as well as digital signal to control them runs over the same line.

Client-Server communication and HD live video feed (&lt;100 ms delay) is achieved between two Linux machines (RPi) using Go programming language and a bit of Shell scripting. Tested successfully in a laboratory environment.

<h2>Requirements</h2>


Required for TCP communication (enough for testing purposes):



1. Raspberry Pi x 2 (Tested on RPi 3 B+)
2. RPi Camera module
3. Ethernet cable, HDMI Display, Keyboard, Mouse
4. Some software (Raspbian Linux, Go programming language, MPlayer… more on that later)

Required for PLC system:



1. Power Line Communications module x 2
2. Power source, Power cable, Filter

Note: If you want to work with PLC, some knowledge in the field of Electronics is required (Working with Power source, some rudimentary knowledge in filters, lab equipment…).

As always, please be extra cautious when you work with electrical current. **You are responsible for yourself and your equipment!** I am not responsible for any damage that may be caused to you or your equipment.

<h2>Hardware setup</h2>


To test the connection and our software we can avoid PLC setup for now and just go for simple connection using Ethernet cable.

Connect two RPi devices using Ethernet cable in their respectful RJ-45 ports.

**( [Client] RPi + HDMI Display, Keyboard, Mouse )&lt;----- Ethernet cable ----->( [Server] RPi + PiCam )**

Attach LEDs on pin 18, 19 and 26 of Server Pi.

Warning: Be careful with LEDs on GPIO pins! (hint: low milliamps output on those)

Client will be set up on the controlling station and will be used to send commands to the Server.

The idea is that the Server will be mounted on the ROV, receive commands from Client to control ROV motors and send live camera feed back to the client.

<h2>Software</h2>


Set up the basic **Raspbian Linux** on both systems, in _raspi-config_ enable **SSH** and **RPi Camera**.  In Raspbian install **Go programming language** and **MPlayer**. In case you need some multimedia codecs (that usually come with Raspbian), feel free to install those as we will be using H.264 for video coding/decoding.

Note: At the time of testing hardware accelerated **OMXPlayer** wasn’t working properly on Raspbian, so MPlayer was used.

Apart from standard libraries, some additional libraries for _Go programming language_ will be needed as well. They can be found in source code (on top of both client and server’s _main.go_ files under _import_ section).

You will need to copy the respected files on your RPi’s, build them and run the binaries on both Client and Server.

Note: To use GPIO pins for generating low frequency software PWM for ESC motor control (or LED dimming), you will need to run Go program using _sudo_ on the Server (or add your user to the appropriate group).

Code in this repo is very simple and is just for testing purposes that you can build upon. \
Ensures the following:



*   Establishes Client-Server communication over TCP.
*   Uses GPIO on Server to generate PWM.
*   Opens and closes H.264 HD video stream with very low latency (&lt;100 ms. Please be aware that it requires a few seconds for proper sync).

<h2>Run</h2>


**Step 1**: Connect all of the hardware (RPi’s, PiCam, Ethernet cable, LEDs on pin 18, 19 and 26)

**Step 2**: Check Go code and Bash scripts for some hard-coded IP addresses as you will probably need to update those. If you don’t know how to find your headless Server IP, use _arp-scan_.

**Step 3**: Build binaries and run (Client on client Pi, Server on server Pi respectfully)

**Step 4**: Establish connection using Client Pi’s screen, keyboard and mouse

**Step 5**: Send test commands:

	R - Dimming of LED on pin 26

	W - Turn on LED on pin 18 and off LED on pin 19

	S - Turn off LED on pin 18 and on LED on pin 19

	A - Turn off both LEDs at pins 18 and 19

	P - Start HD video stream

	O - Stop HD video stream

Note: Ideally, video transmission should be achieved over UDP. Unfortunately there were too many issues with drivers/codec/player in Raspbian Linux at the time. TCP approach works more than good enough as the video is crisp, HD quality and lag free.
