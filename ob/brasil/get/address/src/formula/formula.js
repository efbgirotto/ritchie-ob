const clc = require("cli-color")
const got = require("got")

async function Run(zipcode) {

    if (true !== /^\d{5}-\d{3}$/.test(zipcode)) {
        console.log(clc.red("Invalid zipcode format"))
        return
    }

    const zipcodeParam = zipcode.replace('-', '')
    const serviceUrl = `https://viacep.com.br/ws/${zipcodeParam}/json/`

    // const { body } = await got(serviceUrl)

    got(serviceUrl).then(({ body }) => {
        console.log(clc.green(body));
    })
}

const formula = Run
module.exports = formula
