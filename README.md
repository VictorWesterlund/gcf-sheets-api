# Google Cloud Functions Sheets API
Get values from any public spreadsheet created with Google Sheets as JSON. 
This program is designed for [Google Cloud Functions](https://cloud.google.com/functions) and requires very minimal set-up.
But it can also be run from the command line as a [stand-alone server](#local-installation) for testing.

### Generate an API key
Open the [API & Services Credentials](https://console.cloud.google.com/apis/credentials) page on Google Cloud and generate a new API key with the "Create Credentials" button. Copy the generated key.

## ☁️ Installation on Google Cloud
You need to have a Google Cloud account with Cloud Functions and the [Google Sheets API](https://console.cloud.google.com/marketplace/product/google/sheets.googleapis.com) enabled to use this program.
All you have to do is click "Enable" and wait a few minutes.

1. **Create a new Cloud Function**
   <br>Open the [Cloud Functions](https://console.cloud.google.com/functions/list) dashboard and create a new function with the "Create Function" button. But before you click next..
2. **Add runtime variables**
   <br>This is where you add your app settings. You need to add 3 runtime variables with the "Add Variable" button at the bottom of the page.
   
   Name|Value
   --|--
   `API_KEY`|The API key generated at the start of this guide
   `SHEET_ID`|ID of the spreadsheet.<br>[*Here's how to find it*](https://developers.google.com/sheets/api/guides/concepts)
   `SHEET_RANGE`|[A1 notation](https://developers.google.com/sheets/api/guides/concepts#expandable-1) of the cells to export with this API
3. **Upload source code**
    <br>[**Download a Cloud Functions compatible ZIP**](https://github.com/VictorWesterlund/gcf-sheets-api/releases) of this program from the releases page. Select "ZIP Upload" under the "Source code" dropdown in the function creator
4. **Deploy**
   <br>Click "Deploy" and wait for the Cloud Function to start up. You should see the URL it created for your function under the function details page. The spreadsheet cells should now show up as JSON.
   
## 🖥️ Local installation
You can also run this program locally as a stand-alone server. There are currently no pre-built binaries of this program, but you can download the Go compiler and run it from the CLI.

1. **Clone this repo**
   ```
   git clone https://github.com/VictorWesterlund/gcf-sheets-api.git
   ```

2. **Download Go**
   <br>[Follow the instructions on go.dev](https://go.dev/dl/). Or download it from a package manager
   ```
   sudo apt-get install golang-go
   ```
   
3. **Environment variables**
   <br>Copy the hidden `.env.example` file to `.env` and set the following values. Save and close when you're done
   
   Name|Value
   --|--
   `API_KEY`|The API key generated at the start of this guide
   `SHEET_ID`|ID of the spreadsheet.<br>[*Here's how to find it*](https://developers.google.com/sheets/api/guides/concepts)
   `SHEET_RANGE`|[A1 notation](https://developers.google.com/sheets/api/guides/concepts#expandable-1) of the cells to export with this API
   
4. **Start the server**
   <br>Start the web server with Go run
   
   ```
   go run .
   ```
   
   > **Note**: The server will listen for HTTP (not HTTPS) connections on localhost:8090
