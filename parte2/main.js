const csv = require('csvtojson');
const path = require('path');
const csvFilePath = path.join(__dirname, '/file.csv');

async function main() {    
    const jsonArray = await csv().fromFile(csvFilePath);
    const map = new Map();
    const results = [];
    const usersArrayTotal = [];
    jsonArray.forEach(currentItem => {
        map.set(currentItem.organizacion, {
            organizacion: currentItem.organizacion
        })
    });

    map.forEach((values)=>{
        usersArrayTotal.push(values);
    })

    usersArrayTotal.forEach((currentItem)=>{
        const filteredUsers = jsonArray.filter(item => (item.organizacion == currentItem.organizacion && item.name == currentItem.name));
        currentItem.users = [];
        const usersCurrentItem = new Map();
        filteredUsers.forEach(currentItem => {
            usersCurrentItem.set(currentItem.usuario, {
                username: currentItem.usuario,
                roles: filteredUsers.map(currentuserRol => currentuserRol.rol)
            })
        });
        usersCurrentItem.forEach((values)=>{
            currentItem.users.push(values);
        })
    })
    console.log("Results: -------------");
    console.log(JSON.stringify(usersArrayTotal))
}

main();

