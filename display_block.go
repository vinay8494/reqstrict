package main

const displayBlockedMessage = `
<!Doctype html>
<title>Proxy prevented this page from loading</title>
<style>
  body { padding: 200px; font: 20px Helvetica}
  proxy { display: block; text-align: center; width: 800px; margin: 0 auto; color: #555; }
</style>

<proxy>
  <p>Proxy prevented this page from loading...</p>
</proxy>
`
