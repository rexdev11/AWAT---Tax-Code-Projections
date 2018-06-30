import puppeteer = require("puppeteer");
import {TEMP_CREDS} from "../.local.env";
import * as fs from "fs";

export async function sleep(time: number) {
    return await setTimeout(() => {}, time);
}

export async function runPup(searchTerm: string): Promise<void> {
    const politicalAdURL: string = `https://www.facebook.com/politicalcontentads/?active_status=all&ad_type=ads-with-political-content&q=${searchTerm}`;
    const Browser = await puppeteer.launch({
        headless: false,
        devtools: true
    });
    const Page = await Browser.newPage();
    await Page.goto(politicalAdURL);
    await Page.type('#' + Object.keys(TEMP_CREDS)[0], TEMP_CREDS.email);
    await Page.type('#' + Object.keys(TEMP_CREDS)[1], TEMP_CREDS.pass);
    const result = await Page.evaluate('new XMLSerializer().serializeToString(document.doctype) + document.documentElement.outerHTML;');
    if (result) {
        console.log('here!');
        fs.writeFileSync('./newdoc.html', result);
    }
    await Page.click('#loginbutton');
    await sleep(5000);
    console.log('after', result);
    return await Promise.resolve();
}

runPup('Jack');