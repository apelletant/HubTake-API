# HubTake-API
[![Go Report Card](https://goreportcard.com/badge/github.com/yabou/HubTake-API)](https://goreportcard.com/report/github.com/yabou/HubTake-API)


All the documention for the packages used, are available on their github repository:
<ul>
  <li><a href="github.com/julienschmidt/httprouter">HttpRouter</a></li>
  <li><a href="github.com/cihub/seelog">SeeLog</a></li>
  <li><a href="github.com/jinzhu/gorm">goRM</a></li>
</ul>

# Routes
<h2>User</h2>
<h3>POST</h3>
<ul><li>"/v1/users"
<p>
  <h4>Requiered Body:</h4>
  </p>
  <p>
  {<br/>
	  &emsp;"UserFirstName" : "foo", <br/>
	  &emsp;"UserLastName" : "bar",<br/>
	  &emsp;"UserEmail" : "foo.bar@foobar.com"<br/>
	  &emsp;"UserPromo": 2021<br/>
  }<br/>
  </p>
</li></ul>
<ul><li>"/v1/take"
<p>
  <h4>Requiered Body:</h4>
  </p>
  <p>
  {<br/>
	  &emsp;"UserEmail" : "foo.bar@foobar.com"<br/>
	  &emsp;"ObjectName" : "screwdriver",<br/>
  }<br/>
  </p>
</li></ul>
  <ul><li>"/v1/return"
  <h4>Requiered Body:</h4>
  </p>
  <p>
  {<br/>
	  &emsp;"UserEmail" : "foo.bar@foobar.com"<br/>
	  &emsp;"ObjectName" : "screwdriver",<br/>
  }<br/>
  </p>
</li></ul>
  
<h3>GET</h3>
<ul><li>"/v1/users"
<p>Return a list of User of type:</p>
  <p>
  [{<br/>
	 &emsp;"UserId": 0</br>
  &emsp;"UserFirstName": "Foo",</br>
  &emsp;"UserLastName": "Bar",</br>
  &emsp;"UserEmail": "foo.bar@foobar.com",</br>
  &emsp;"UserObjectId": null,</br>
  &emsp;"UserHasObject": 0,</br>
  &emsp;"UserPromo": 2012,</br>
  }]
  </p>
</li></ul>
<ul><li>"/v1/users/:userEmailGet"
<h4>Requiered Parameter:</h4>
Email addres of the user you're loocking for</br>
Return a User object of type:</br>
  <p>
  [{<br/>
	 &emsp;"UserId": 0</br>
  &emsp;"UserFirstName": "Foo",</br>
  &emsp;"UserLastName": "Bar",</br>
  &emsp;"UserEmail": "foo.bar@foobar.com",</br>
  &emsp;"UserObjectId": null,</br>
  &emsp;"UserHasObject": 0,</br>
  &emsp;"UserPromo": 2012,</br>
  }]<br/>
  </p>
</li></ul>
<ul><li>"/v1/user/byID/:userId"
<h4>Requiered Parameter:</h4>
ID of the user you're loocking for</br>
Return a User object of type:</br>
  <p>
  [{<br/>
	 &emsp;"UserId": 0</br>
  &emsp;"UserFirstName": "Foo",</br>
  &emsp;"UserLastName": "Bar",</br>
  &emsp;"UserEmail": "foo.bar@foobar.com",</br>
  &emsp;"UserObjectId": null,</br>
  &emsp;"UserHasObject": 0,</br>
  &emsp;"UserPromo": 2012,</br>
  }]<br/>
  </p>
</li></ul>
<ul><li>"/v1/usersHasObject"
<p>Return a list of User whom borrow an object, of type:</p>
  <p>
  [{<br/>
	 &emsp;"UserId": 0</br>
  &emsp;"UserFirstName": "Foo",</br>
  &emsp;"UserLastName": "Bar",</br>
  &emsp;"UserEmail": "foo.bar@foobar.com",</br>
  &emsp;"UserObjectId": null,</br>
  &emsp;"UserHasObject": 0,</br>
  &emsp;"UserPromo": 2012,</br>
  }]<br/>
  </p>
</li></ul>
<h3>DELETE</h3>
<ul><li>"/v1/users/:userToDelete"
<h4>Requiered Parameter:</h4>
Email address of the user you want to delete.</br>
Return data of the deleted user:<br/>
  <p>
  [{<br/>
	 &emsp;"UserId": 0</br>
  &emsp;"UserFirstName": "Foo",</br>
  &emsp;"UserLastName": "Bar",</br>
  &emsp;"UserEmail": "foo.bar@foobar.com",</br>
  &emsp;"UserObjectId": null,</br>
  &emsp;"UserHasObject": 0,</br>
  &emsp;"UserPromo": 2012,</br>
  }]<br/>
  </p>
</li></ul>
<h2>Object</h2>
<h3>POST</h3>
<ul><li>"/v1/objects/post/:objectName"
<p>
  <h4>Requiered Parameter:</h4>
  Name of the object to be added.</br>
</p>
<p>
    <h4>Requiered Parameter:</h4>
  On success return 204
  </p>
</li></ul>
<h3>GET</h3>
<ul><li>"/v1/objects"
<p>Return a list of every Object, of type:</p>
  <p>
  [{<br/>
    &emsp;"ObjectID": 1,<br/>
    &emsp;"ObjectName": "screwdriver",<br/>
    &emsp;"ObjectIsTaken": 0,<br/>
    &emsp;"ObjectDateBorrow": 0,<br/>
    &emsp;"ObjectDateReturn" 0:<br/>
  }]<br/>
  </p>
</li></ul>
<ul><li>"/v1/objects/isTaken"
<p>Return a list of every borrowed Object, of type:</p>
  <p>
  [{<br/>
    &emsp;"ObjectID": 1,<br/>
    &emsp;"ObjectName": "screwdriver",<br/>
    &emsp;"ObjectIsTaken": 0,<br/>
    &emsp;"ObjectDateBorrow": 0,<br/>
    &emsp;"ObjectDateReturn" 0:<br/>
  }]<br/>
  </p>
</li></ul>

<ul><li>"/v1/objects/notTaken"
<p>Return a list of every free to take Object, of type:</p>
  <p>
  [{<br/>
    &emsp;"ObjectID": 1,<br/>
    &emsp;"ObjectName": "screwdriver",<br/>
    &emsp;"ObjectIsTaken": 0,<br/>
    &emsp;"ObjectDateBorrow": 0,<br/>
    &emsp;"ObjectDateReturn" 0:<br/>
  }]<br/>
  </p>
</li></ul>

<ul><li>"/v1/objects/getByName/:name"
    <p>Return the object wanted</p>
    <h4>Requiered Parameter:</h4>
The name of the object.
<p>Return the Object's data, of type:</p>
  <p>
  {<br/>
    &emsp;"ObjectID": 1,<br/>
    &emsp;"ObjectName": "screwdriver",<br/>
    &emsp;"ObjectIsTaken": 0,<br/>
    &emsp;"ObjectDateBorrow": 0,<br/>
    &emsp;"ObjectDateReturn" 0:<br/>
  }<br/>
  </p>
</li></ul>
<h3>DELETE</h3>
<ul><li>"/v1/objects/:objectToDelete"
    <p>Delete the object wanted</p>
<h4>Requiered Parameter:</h4>
The name of the object.<br/>
On success return 204
</li></ul>

<h2>Plastic</h2>
<h3>POST</h3>
<ul><li>"/v1/plastic"
    <p>Add a platic</p>
    <h4>Requiered body:</h4>
    <p>
        {<br/>
          &emsp;"PlasticColor": "orange",<br/>
          &emsp;"PlasticPrice": 12<br/>
        }<br/>
    </p>
</li></ul>
</br>

<h3>GET</h3>
<ul><li>"/v1/plastic"
    <p>Return all the plastic available</p>
    <h4>Response:</h4>
    <p>
        [{<br/>
          &emsp;"PlasticID" : 0,<br/>
          &emsp;"PlasticColor" :"Orange",<br/>
          &emsp;"PlasticPrice" : 12<br/>
        }]<br/>
    </p>
</li></ul>
</br>
<ul><li>"/v1/plastic/:plasticColor"
    <p>Get plastic data of the color given</p>
    <p><h4>Requiered parameter:</h4>
    color wanted</p>
    <p>Response body:</p>
    <p>
      {<br/>
      &emsp;"PlasticID" : 0,<br/>
      &emsp;"PlasticColor" :"Orange",<br/>
      &emsp;"PlasticPrice" : 12<br/>
      }<br/>
  </p>
</li></ul>
</br>

<h2>Command</h2>
<h3>POST</h3>
<ul><li>
  "/v1/command"
  <h4>Add a command</h4>
  <h5>Requiered body:</h5>
  <p>
      {<br/>
        &emsp;"Email" :"foo.bar@foobar.com",<br/>
        &emsp;"PlasticColor" :"Orange",<br/>
        &emsp;"URLModel" :"mode.url",<br/>
        &emsp;"Length" :"123"<br/>
      }<br/>
  </p>
</li></ul>
</br>

<h3>GET</h3>	
<ul><li>"/v1/command"
    <p>Return all the command</p>
    <h4>Response:</h4>
    <p>
        [{<br/>
          &emsp;"CommandID" : 0,<br/>
          &emsp;"CommandModel" : "modelUrl",<br/>
          &emsp;"CommandIDUser" : 1567,<br/>
          &emsp;"CommandIDPlastic" : 173,<br/>
          &emsp;"CommandPrice" : 7,<br/>
          &emsp;"CommandLength" : 12<br/>
        }]<br/>
    </p>
</li></ul>
</br>

<h3>DELETE</h3>
<ul><li>"/v1/command/:idCommand"
    <p>Delete the command of the given id</p>
    <h4>Response:</h4>
    <p>204 on succes
    </p>
</li></ul>
</br>
