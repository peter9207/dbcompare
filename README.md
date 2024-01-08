# Trying to benchmark databases
This project is my attempt at understanding database performance. 

While interviewing, I'm often asked to give reasons on database selections based on usage requirements, but the criterias are incomplete. Making decisions like this in real life requires us to think about use cases, schema, amount of data, traffic patterns, and other functional and non-functional requirements. So I often have trouble justifying the decisions. 

So my thinking is to create a repo that has a simply schema, and write workers that perform simple workflows, such as those in CRUD endpoints to measure how many reads a certain type of database can handle. 


