# Go Practical Programming Tutoral 13 Revised

Revised code for [this tutorial](https://www.youtube.com/watch?v=dySAX8VZ2TU)

## Problems Solved:
1. URL for Washington Post sitemap has been changed from "https://wwww.washingtonpost.com/news-sitemap-index.xml" 
   (given in the tutorial) to the new URL https://www.washingtonpost.com/news-sitemaps/index.xml.  I have changed the 
   url to reflect its current form.  If it has changed again, feel free and flag an issue in the repo.
1. You will find that ```http.get("https://www.washingtonpost.com/news-sitemaps/index.xml")``` is extremely slow.   
   I struggled with this untilI ran across the solution on 
   stack overflow:  Washingtonpost.com now has anti bot scripts that slow down traffic from devices flagged as bots.  
   If you use ```http.get()``` by itself with no headers, washingtonpost.com immidiately flags you as a bot and slows you 
   down.  You need to "trick" the site into thinking your not a bot by creating a custom get request with standard 
   looking headings.  I ended up making it a separate method so that I could call it multiple times without rewriting a 
   bunch of duplicate code.  Definitely going to speed up that for loop for the sub sitemaps a whole lot.
   The revised script is as follows:
   ```Go
   client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.washingtonpost.com/news-sitemaps/index.xml", nil)
	req.Header.Set("Connection","Keep-Alive")
	req.Header.Set("Accept-Language","en-US")
	req.Header.Set("User-Agent","Mozilla/5.0")
	resp, _ := client.Do(req)
	bytes, _ := ioutil.ReadAll(resp.Body)
   ```