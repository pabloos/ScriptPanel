## p2p tests

In order to declouping the test enviorement I'm using Docker containers:

- One for a Node.js running Jest and Puppeteer

- The other run a [Browserless](https://github.com/joelgriffith/browserless) Chrome that will connect to our webserver

Both containers stands in a diferent network. In order to connect sp-net and test-net there is also a third network wich allows the connection between the browserless server and the webserver.
