# inventory

### What is this repository for?

This is frontend for inventory management system

### How do I get set up?

- Install dependencies

  `yarn install`

- Build project

  `yarn build`

- Copy build/ into release/ and run
  `docker build . -t inventory-ui`
  `docker run -d --publish 15888:15888 --name docker-inventory-ui inventory-ui`
