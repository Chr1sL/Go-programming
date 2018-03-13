this is the pink slip project folder

So here's the link https://golang.org/doc/articles/wiki/ to the  stuff that makes a site

General structure options:
 1. we can write the entire thing in go and work with that
 2. we can try to have it that go does server sides stuff (error handling,  getting information, saving, and sending) and then JS and css handles the looks sending stuff to the go files [I'm thinking of having an intermediary file where the JS saves the inputed information, and then go finds it and passes the information on to molly and the teacher involved ] 
 
 https://golang.org/pkg/net/smtp/  << that is a link to help with automatic email sending.
 
 https://stackoverflow.com/questions/23282311/parse-input-from-html-form-in-golang << will help us when we need to be able to get the data from the forms to send along
 
 I am thinking that hte flow of this program should be something like: 
 *Build the web page (it doesnt hafe to be pretty or pink wee can fix that later)
 *collect the information (name, teacher, class, block, teacher email, & date/time can be autofilledusing the time package)
 *send the information to both Molly and to the teacher specified by the email provided

just a note: once we finish the code we will need a server to to host the page but we can ask the school or something we'll cross that bridge when we get to its
