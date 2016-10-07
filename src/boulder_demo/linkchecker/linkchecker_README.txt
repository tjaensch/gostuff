+++ XML LINKCHECKER COMMAND LINE TOOL +++

WHAT DOES IT DO?
Check all http:// and https:// links in XML files in directory the tool is run in, report back the server responses received from the checked links and log the failed responses (everything other than "200 OK") to a linkchecker_bad_links_log file in the current working directory

PREREQUISITES
The program uses Bash and cURL commands under the hood and will not work if Bash and/or cURL are not installed in the environment it's run in

- Check on Linux:  Run "curl -V" and "bash --version" to see if they're installed
- Check on Mac: Run "curl -V" and "bash --version" to see if they're installed
- Windows: Not supported

HOW TO RUN IT?
- Drop the binary suitable for your system into the folder with the XML files you want to check for broken links
- Open a shell (PuTTY, Mac terminal, whatever) and navigate to the folder where you dropped the linkchecker binary
- Run "chmod -R 777 linkchecker"
- Run "./linkchecker" and if it works you should be able to see what it's doing in your shell and end up with a linkchecker_bad_links_log file in the current working directory after it's done

QUESTIONS/IT'S NOT WORKING...!?
thomas.jaensch@noaa.gov

