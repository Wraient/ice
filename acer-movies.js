<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Acer Movies</title>
    <link rel="stylesheet" href="/main.css">
</head>
<style>
body {
  scrollbar-width: thin;
  scrollbar-color: #aaa transparent; /* Adjust colors as needed */
}

body::-webkit-scrollbar {
  width: 8px;
  height: 8px; /* Important for consistent width and height */
}

body::-webkit-scrollbar-track {
  background: transparent; /* or a light gray for a slightly different look */
}

body::-webkit-scrollbar-thumb {
  background-color: #aaa; /* Adjust thumb color as needed */
  border-radius: 4px; /* Optional: Rounded corners for a smoother look */
}
  
.loader-box{
    display: flex;
    justify-content: center;
    width: 100%;
    height: 100%;
}
.loader {
  width: 16px;
  height: 16px;
  padding: 4px;
  margin: 4px;
  aspect-ratio: 1;
  border-radius: 50%;
  background: #ffffff;
  --_m: 
    conic-gradient(#0000 10%,#000),
    linear-gradient(#000 0 0) content-box;
  -webkit-mask: var(--_m);
          mask: var(--_m);
  -webkit-mask-composite: source-out;
          mask-composite: subtract;
  animation: l3 1s infinite linear;
}
@keyframes l3 {to{transform: rotate(1turn)}}
</style>
<body class="bg-neutral text-neutral-content">
    <div class="app">

        <!-- Nav Bar -->
       <nav class="lg:hidden p-6 flex justify-between items-center border-b border-base">
            <span class="text-xl">
                <a class="content-center" href="/">
                    Acer Movies
                </a>
            </span>

            <div class="flex">
                <a class="p-2 hover:border-primary mr-4  text-xs text-primary border border-base rounded-lg" href="https://discord.gg/hpkMqrwRCE" target="_blank" >
                    <div class="flex ">
                        <svg class="w-6 h-6" viewBox="0 0 192 192" xmlns="http://www.w3.org/2000/svg" fill="none"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><path stroke="#e01b24" stroke-linecap="round" stroke-linejoin="round" stroke-width="12" d="m68 138-8 16c-10.19-4.246-20.742-8.492-31.96-15.8-3.912-2.549-6.284-6.88-6.378-11.548-.488-23.964 5.134-48.056 19.369-73.528 1.863-3.334 4.967-5.778 8.567-7.056C58.186 43.02 64.016 40.664 74 39l6 11s6-2 16-2 16 2 16 2l6-11c9.984 1.664 15.814 4.02 24.402 7.068 3.6 1.278 6.704 3.722 8.567 7.056 14.235 25.472 19.857 49.564 19.37 73.528-.095 4.668-2.467 8.999-6.379 11.548-11.218 7.308-21.769 11.554-31.96 15.8l-8-16m-68-8s20 10 40 10 40-10 40-10"></path><ellipse cx="71" cy="101" fill="#e01b24" rx="13" ry="15"></ellipse><ellipse cx="121" cy="101" fill="#e01b24" rx="13" ry="15"></ellipse></g></svg>
                    
                    </div>
                </a>

                <a class="p-2 hover:border-primary text-xs text-primary border border-base rounded-lg" href="/contact">
                    <div class="flex ">
                        <svg class="w-6 h-6" viewBox="0 0 128 128" version="1.1" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" fill="#000000"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <style type="text/css"> .st0{display:none;} .st1{display:inline;} .st2{fill:none;stroke:#e01b24;stroke-width:8;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;} </style> <g class="st0" id="Layer_1"></g> <g id="Layer_2"> <path class="st2" d="M20.1,40l38.1,31.5c3.3,2.8,8.2,2.8,11.5,0L107.9,40"></path> <path class="st2" d="M107.9,85.6V36.1c0-2.1-1.7-3.7-3.7-3.7H23.9c-2.1,0-3.7,1.7-3.7,3.7v55.7c0,2.1,1.7,3.7,3.7,3.7h84"></path> </g> </g></svg>
                    
                    </div>
                </a>

            </div>
           
        </nav>

        <!-- Search Bar -->
        <div class=" p-6 sticky top-0 z-5 border-b border-base bg-neutral flex justify-center lg:justify-between ">
            <span class="hidden lg:flex text-xl">
                <a class="content-center" href="/">
                    Acer Movies
                </a>
            </span>

            <div class="bg-base-200 flex rounded border border-base w-full max-w-160">
                <input class="w-full p-2 px-4 text-sm outline-none" id="searchInput" placeholder="Search something here...">
                <button class="p-2 cursor-pointer bg-primary disabled:bg-base-300 rounded border-l border-base" id="searchButton">
                    <svg class="w-6 h-6 " aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                        <path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21-3.5-3.5M17 10a7 7 0 1 1-14 0 7 7 0 0 1 14 0Z"/>
                      </svg>
                </button>
            </div>

            <div class="hidden lg:flex">
                <a class="p-2 hover:border-primary mr-4  text-xs text-primary border border-base rounded-lg" href="https://discord.gg/hpkMqrwRCE" target="_blank" >
                    <div class="flex ">
                        <svg class="w-6 h-6" viewBox="0 0 192 192" xmlns="http://www.w3.org/2000/svg" fill="none"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><path stroke="#e01b24" stroke-linecap="round" stroke-linejoin="round" stroke-width="12" d="m68 138-8 16c-10.19-4.246-20.742-8.492-31.96-15.8-3.912-2.549-6.284-6.88-6.378-11.548-.488-23.964 5.134-48.056 19.369-73.528 1.863-3.334 4.967-5.778 8.567-7.056C58.186 43.02 64.016 40.664 74 39l6 11s6-2 16-2 16 2 16 2l6-11c9.984 1.664 15.814 4.02 24.402 7.068 3.6 1.278 6.704 3.722 8.567 7.056 14.235 25.472 19.857 49.564 19.37 73.528-.095 4.668-2.467 8.999-6.379 11.548-11.218 7.308-21.769 11.554-31.96 15.8l-8-16m-68-8s20 10 40 10 40-10 40-10"></path><ellipse cx="71" cy="101" fill="#e01b24" rx="13" ry="15"></ellipse><ellipse cx="121" cy="101" fill="#e01b24" rx="13" ry="15"></ellipse></g></svg>
                    
                    </div>
                </a>

                <a class="p-2 hover:border-primary text-xs text-primary border border-base rounded-lg" href="/contact">
                    <div class="flex ">
                        <svg class="w-6 h-6" viewBox="0 0 128 128" version="1.1" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" fill="#000000"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <style type="text/css"> .st0{display:none;} .st1{display:inline;} .st2{fill:none;stroke:#e01b24;stroke-width:8;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:10;} </style> <g class="st0" id="Layer_1"></g> <g id="Layer_2"> <path class="st2" d="M20.1,40l38.1,31.5c3.3,2.8,8.2,2.8,11.5,0L107.9,40"></path> <path class="st2" d="M107.9,85.6V36.1c0-2.1-1.7-3.7-3.7-3.7H23.9c-2.1,0-3.7,1.7-3.7,3.7v55.7c0,2.1,1.7,3.7,3.7,3.7h84"></path> </g> </g></svg>
                    
                    </div>
                </a>

            </div>

        </div>

        <!-- Alert Toast Box -->
        <div id="alertToastBox" class="fixed z-10 top-20 left-0 w-full pointer-events-none flex flex-col items-center ">

            
        </div>

        <!-- <div class="p-6 py-4 flex w-full justify-center border-b border-base cursor-pointer" onclick="this.style.display='none'">
            <div class="my-2 p-4 text-sm  rounded-lg border border-primary">
                The server issue has been temporarily patched.
            </div>
        </div> -->


        <!-- Search Result Box -->
        <div class="min-h-screen ">
            <div class="py-6 flex flex-wrap" id="searchResultBox">

                    <div class="py-4 px-6 flex items-center">
                        <span class="text-2xl mr-4 ">Trending</span>

                    </div>

                    <div id="trendingBox" class="flex overflow-x-scroll w-full">
                        

                    </div>

                    <div class="py-6  px-6 my-2 flex items-center">
                        <span class="text-2xl mr-4 ">History</span>

                        <label class="inline-flex items-center cursor-pointer">
                            <input type="checkbox" id="historyCheckbox" value="" checked class="sr-only peer">
                            <div class="relative w-9 h-5 bg-base peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-gray-700  rounded-full peer  peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all  peer-checked:bg-primary "></div>
                            <span class="ms-3 text-xs font-medium text-gray-300 ">Disabling will clear the history state</span>
                        </label>
  
                    </div>

                    <div id="historyBox" class="flex overflow-x-scroll w-full">
                        

                    </div>
                </div>

            </div>
            
        </div>
        


        <!-- Footer -->
        <div class="p-6 border-t border-base flex justify-between flex-wrap text-sm  ">
            <a class="mx-4 my-2 underline" href="/terms">Terms of Service</a>
            <a class="mx-4 my-2 underline" href="/dmca">DMCA</a>
            <a class="mx-4 my-2 underline" href="/contact">Contact Us</a>
            <span id="showStats" class="mx-4 my-2">© Acer Movies. All rights reserved</span>

        </div>
        <div id="userStats" class="p-6 hidden border-t border-base flex justify-between flex-wrap text-sm  ">
               
        </div>


    </div>
    
</body>
<script>

  

let searchResultBox = document.querySelector("#searchResultBox")
let historyBox = document.querySelector("#historyBox")


const historyCheckbox = document.querySelector("#historyCheckbox")


function saveContentState(data){
    history.pushState( data, null , window.location.href)
    
}

window.addEventListener('popstate',(e)=>{

    //add
    let type = e.state?.type ?? null
    let data = e.state ?? null

    if ( type === "searchResults") {
        createResultPage(data.searchResult , data.searchQuery)
    }else if (type === "qualityPage") {
        createQualityPage( data.source , data.qualityList )
    }else if (type === "episodesPage") {
        createEpisodesPage( data.source , data.episodesList )
    }else{
        clearResultBox()
    }
   

        

    
})


async function postRequest( url , data ){

    return await fetch(url, {
        method: 'POST', 
        headers: {
            'Content-Type': 'application/json' 
        },
        body: JSON.stringify(data) 
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
    })
    .then(responseData => {
        console.log('Success:', responseData);
        return responseData
    })
    .catch(error => {
        console.error('Error:', error);
    });
}


document.querySelector("#searchInput").addEventListener("keydown",(e)=>{
    if(e.keyCode == 13) {
        document.querySelector("#searchButton").click()
    }
  })

function clearResultBox(){
    // remove all elements
    while (searchResultBox.firstChild) {
        searchResultBox.firstChild.remove()
    }
}


function createResultPage( searchResult , searchQuery ){

    clearResultBox()

    searchResult.forEach(data => {
        let div = document.createElement("div")
        div.innerHTML = `
            <img class="h-32 rounded-lg" src="${data.image}">
            <span class="p-4 text-sm">
                ${data.title}
            </span>
        `
        
        div.classList.add("cursor-pointer","basis-md","flex-1","bg-base-300","w-full", "m-4", "flex", "rounded-lg", "border", "border-base", "hover:border-primary",)
        div.addEventListener('click',()=>{
            getSourceQuality(data)
        })

        searchResultBox.appendChild(div)
    })


    
}



document.querySelector("#searchButton").addEventListener("click",async function(){

    searchResultBox.innerHTML = '<div class="loader-box">Searching... <div class="loader"></div></div>'
    this.disabled = true
    let searchQuery = document.querySelector("#searchInput")?.value || ""
    let data = {
        searchQuery : searchQuery
    }

    let responseData = await postRequest( '/api/search' , data )


    if (responseData?.searchResult?.length > 0) {
        
        createResultPage( responseData.searchResult , searchQuery)
        saveContentState( { 
            searchResult :  responseData.searchResult ,
            searchQuery : searchQuery ,
            type : "searchResults" 
        }  )

    }else{
        searchResultBox.innerHTML = `
        <div class="px-6 flex flex-col w-full jusitfy-center items-center ">
            <span class=" underline text-2xl my-4"> No result found for "${searchQuery}" </span>
            
            <div class="my-2 p-4 text-sm  rounded-lg border border-base">
                <strong>Use Official Movie Titles:</strong> Search using the official name. Ex: searching for "Spider-Man" will yield results, while "Spiderman" might not.
            </div>
            <div class="my-2 p-4 text-sm  rounded-lg border border-base">
                <strong>Try Partial Keywords:</strong> If you're unsure of the full title, use partial words like "Spider" or "Spider-Man" to help find the movie you're looking for.
            </div>
            <div class="my-2 p-4 text-sm  rounded-lg border border-base">
                <strong>Search by IMDb IDs:</strong> If you know the IMDb ID, entering it will give you precise results. For example, "tt0145487" corresponds to "Spider-Man."
            </div>
        </div>
        `
    }

    this.disabled = false

})

function createQualityPage(source , qualityList){
        
        clearResultBox()

        makeCard(source)
        qualityList.forEach(item => {
          
            if (item.episodesUrl == "https://dramadrip.com/home/") {
                return
            }
            let qualityInfo = item.quality + item.title.split(item.quality)?.[1]

            let div = document.createElement("div")
            // item.episodesUrl ? item.title : qualityInfo / in span to short the title
            div.innerHTML = `
                <a class="p-4 m-2 w-full max-w-md cursor-pointer rounded border border-base bg-base-300 hover:border-primary text-sm flex justify-center items-center">
                    <span>${ item.title }</span>
                    
                     <button class="${item.episodesUrl ? "hidden" : ""} cursor-pointer p-2 ml-4 bg-base disabled:bg-base-300 rounded border-l border-base" id="searchButton">
                        <svg class="w-6 h-6 " aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V4M7 14H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2m-1-5-4 5-4-5m9 8h.01"/>
                        </svg>
                    </button>
                </a>
            `
            div.classList.add("quality-box","basis-md","flex-1", "flex","justify-center","w-full")
            let parent = div.querySelector("a")

            let clickHandler = function() {
                
                if (item.episodesUrl) {
                    getSourceEpisodes(source , item  )
                }else{
                    
                    getSourceUrl(item.url , "movie",this,clickHandler)
                }

            }

            parent.addEventListener('click',clickHandler)

            searchResultBox.appendChild(div)

        });
}


//  get Source Quality
function makeCard(source){
    
    // Add Current Source element
    let div = document.createElement("div")
    div.innerHTML = `
        <div class="basis-md max-w-160 flex-1 bg-base-300 w-full m-4 flex rounded-lg border border-primary">
            <img class="h-32 rounded-lg" src="${source.image}">
            <span class="p-4 text-sm">
                ${source.title}
            </span>
        </div>
    `
    div.classList.add("flex","justify-center","w-full")
    searchResultBox.appendChild(div)

    searchResultBox.innerHTML += '<div class="w-full h-px m-4 bg-base"></div>'

   
}


async function getSourceQuality(source){

    clearResultBox()
    // make Card
    makeCard(source)

    if (historyCheckbox.checked === true) {
        saveHistoryCard(source)  
    }

     // Message User Info feedback
     searchResultBox.innerHTML += '<div class="loader-box">Finding Source Quality... <div class="loader"></div></div>'


    let data = {
        url : source.url
    }
    // Fetching Source Quality
    let responseData = await postRequest( '/api/sourceQuality' , data )


    // creating quality boxes
    if (responseData?.sourceQualityList?.length > 0) {

        createQualityPage(source ,responseData.sourceQualityList )
 
        saveContentState({ 
            type : "qualityPage" ,
            source : source ,
            qualityList : responseData.sourceQualityList 
        })

    }else{
        searchResultBox.innerHTML = '<div class="loader-box">Error! Try Refreshing.</div>'
    }
    
}



function createEpisodesPage(source , episodesList){


    clearResultBox()

    makeCard(source)

    episodesList.forEach(episode => {
        let div = document.createElement("div")
        div.innerHTML = `
            <a class="p-4 m-2 cursor-pointer rounded border border-base bg-base-300 hover:border-primary text-sm flex justify-center items-center">
                <span>${ episode.title }</span>
                
                <button class=" p-2 ml-4 cursor-pointer bg-base disabled:bg-base-300 rounded border-l border-base" id="searchButton">
                    <svg class="w-6 h-6 " aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V4M7 14H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2m-1-5-4 5-4-5m9 8h.01"/>
                    </svg>
                </button>
            </a>
        `
        
        div.classList.add("quality-box","basis-md","flex-1", "flex","justify-center","w-full")
        let parent = div.querySelector("a")

        let clickHandler = function() {
            
            getSourceUrl(episode.link , "episode",this,clickHandler)
            
        }

        parent.addEventListener('click',clickHandler)

        searchResultBox.appendChild(div)

    });
}


async function getSourceEpisodes(source , item){

    clearResultBox()
    // make Card
    makeCard(source)

     // Message User Info feedback
     searchResultBox.innerHTML += '<div class="loader-box">Finding Source Episodes... <div class="loader"></div></div>'

    

    let data = {
        url : item.episodesUrl
    }
    // Fetching Source Quality
    let responseData = await postRequest( '/api/sourceEpisodes' , data )


    // creating quality boxes
    if (responseData?.sourceEpisodes?.length > 0) {

        createEpisodesPage(source , responseData.sourceEpisodes )
        saveContentState({ 
            type : "episodesPage" ,
            source : source ,
            episodesList : responseData.sourceEpisodes 
        })

    }else{
        searchResultBox.innerHTML += " *** No source episodes ***"
    }
    
}



function alertToast(msg){
    let alertBox = document.querySelector("#alertToastBox")
    let alertToast = document.createElement('div')
    alertToast.innerText = msg

    alertToast.classList.add("my-4","animate-bounce", "pointer-events-auto", "px-6", "py-4", "border", "border-primary", "rounded", "shadow-md", "shadow-black", "bg-base-300")

    alertBox.appendChild(alertToast)

    setTimeout(()=>{
        alertBox.removeChild(alertToast)
    },3000)
}



let sourceUrlCounter = 0
let sourceUrlLimit = 3

async function getSourceUrl(url ,seriesType, parent, clickHandler){

    // Check limit
    if (sourceUrlCounter >= sourceUrlLimit) {
        return alertToast(`Please wait, limit exceeded.[${sourceUrlLimit}/${sourceUrlCounter}]`)
    }

    sourceUrlCounter ++

    console.log(parent)

    parent.removeEventListener("click",clickHandler)

    parent.classList.add("border-neutral-content")

    let button = parent.querySelector("button")
    button.innerHTML = '<div class="loader"></div>'

    let data = {
        url : url,
        seriesType : seriesType
    }

// // Fetching Source Quality
let responseData = await postRequest( '/api/sourceUrl' , data )


// creating quality boxes
    if (responseData?.sourceUrl) {

            button.innerHTML = `
                <svg class="w-6 h-6 " aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V4M7 14H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2m-1-5-4 5-4-5m9 8h.01"/>
                </svg>
            `
            button.classList.add("bg-primary")
            
            parent.removeEventListener("click",clickHandler)
            parent.disabled = false
            parent.href = responseData.sourceUrl
            
            
    }else{
        searchResultBox.innerHTML = '<div class="loader-box">Error! Try Refreshing.</div>'
    }

    sourceUrlCounter --
}







function saveHistoryCard(obj) {
    if (!obj || !obj.url) {
      console.error('Object must have a url property.');
      return;
    }
  
    const storageKey = 'historyArray';
  
    let objectsArray = JSON.parse(localStorage.getItem(storageKey)) || [];
  
    // if object already exist then remove it
    const existingIndex = objectsArray.findIndex(item => item.url === obj.url);
    if (existingIndex !== -1) {
      objectsArray.splice(existingIndex, 1);
    }
  
    // Add the new object at the start
    objectsArray.unshift(obj);
  
    //save
    localStorage.setItem(storageKey, JSON.stringify(objectsArray));
  }
  


//  get Source Quality
function makeHistoryCard(source){
    
    // Add Current Source element
    let div = document.createElement("div")
    div.innerHTML = `
        <div class="w-32 max-w-32 cursor-pointer bg-base-300 m-4 flex flex-col rounded-lg border border-base">
            <img class="h-42 object-center object-cover rounded-t-lg" src="${source.image}">
            <span class="p-4 px-2 text-neutral-content text-xs truncate">
                ${source.title}
            </span>
        </div>
    `
    div.classList.add("flex")
    div.addEventListener('click', ()=>{ 
        getSourceQuality(source)
    })
    historyBox.appendChild(div)

}

function loadHistoryCard(){
    let objectsArray = JSON.parse(localStorage.getItem("historyArray")) || [];
    objectsArray.forEach(item=> { makeHistoryCard(item)})

}


historyCheckbox.addEventListener('change', () => {
    localStorage.setItem("historyCheckbox", JSON.stringify(historyCheckbox.checked));
    if (historyCheckbox.checked === true) {
        loadHistoryCard()
    }
    if (historyCheckbox.checked === false) {
        localStorage.setItem("historyArray", JSON.stringify([])); 
        historyBox.innerHTML =""   
    }
});


window.addEventListener('load', () => {
    const storedState = localStorage.getItem("historyCheckbox");
    if (storedState !== null) {
      historyCheckbox.checked = JSON.parse(storedState);
    }

    if (historyCheckbox.checked === true) {
        loadHistoryCard()
      }
});


const trendingBox = document.querySelector("#trendingBox")

//  get Source Quality
function makeTrendingCard(source){
    
    // Add Current Source element
    let div = document.createElement("div")
    div.innerHTML = `
        <div class="w-32 max-w-32 cursor-pointer bg-base-300 m-4 flex flex-col rounded-lg border border-base">
            <img class="h-42 object-center object-cover rounded-t-lg" src="${source.image}">
            <span class="p-4 px-2 text-neutral-content text-xs truncate">
                ${source.title}
            </span>
        </div>
    `
    div.classList.add("flex")
    div.addEventListener('click', ()=>{ 
        getSourceQuality(source)
    })
    trendingBox.appendChild(div)

}

const trendingList = [

    {
		"title": "Wednesday (Season 1-2) Dual Audio {Hindi-English} With Esubs WeB- DL 480p [210MB] || 720p [350MB] || 1080p [1.6GB]",
		"url": "https://moviesmod.cafe/download-wednesday-hindi-480p-720p-1080p/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2022/11/Download-Wednesday-MoviesMod.jpeg"
	},
    {
		"title": "Squid Game (Season 1-3) Multi Audio {Hindi-English-Korean} WeB-DL 480p [220MB] || 720p [330MB] || 1080p [1.5GB]",
		"url": "https://moviesmod.cafe/download-squid-game-hindi-480p-720p-1080p/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2025/06/Download-Squid-Game-MoviesMod.jpg"
	},
    {
		"title": "Jurassic World: Rebirth (2025) Dual Audio {Hindi-English} WEB-DL 480p [500MB] || 720p [1.2GB] || 1080p [2.9GB]",
		"url": "https://moviesmod.cafe/download-jurassic-world-rebirth-2025-hindi-english-480p-720p-1080p/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2025/08/Jurassic-World-Rebirth-2025-MoviesMod.tube_.jpg"
	},
    {
		"title": "Mission: Impossible – The Final Reckoning (2025) Dual Audio (Hindi-English) Msubs IMAX Web-Dl 480p [580MB] || 720p [1.6GB] || 1080p 3.8[GB]",
		"url": "https://moviesmod.cafe/download-mission-impossible-the-final-reckoning-2025-hindi-english-480p-720p-1080p-2/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2025/08/Download-Mission-Impossible-The-Final-Reckoning-MoviesMod.jpeg"
	},
    {
		"title": "F1: The Movie (2025) Dual Audio {Hindi-English} WEB-DL 480p [550MB] || 720p [1.4GB] || 1080p [3.4GB]",
		"url": "https://moviesmod.cafe/download-f1-the-movie-2025-hindi-english-480p-720p-1080p/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2025/06/Download-F1-The-Movie-MoviesMod.jpg"
	},
    {
		"title": "Superman (2025) Dual Audio {Hindi-English} iMAX WEB-DL 480p [490MB] || 720p [1.2GB] || 1080p [2.9GB]",
		"url": "https://moviesmod.cafe/download-superman-2025-hindi-english-480p-720p-1080p/",
		"image": "https://moviesmod.cafe/wp-content/uploads/2025/08/Superman-2025-MoviesMod.cafe_.jpg"
	},




]

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
    
    array.forEach(item=>{
        makeTrendingCard(item)
    })
}
shuffleArray(trendingList)



  
document.querySelector("#showStats").addEventListener('click',()=>{
  document.querySelector("#userStats").style.display="flex"
})

  fetch('/main/stats')
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return response.json();
  })
  .then(data => {
    let list = data.daysCounter
    for (const key in list) {
      if (list.hasOwnProperty(key)) {
        addUserStat(`${key} : ${list[key]}`)
      }
    }
  })
  .catch(error => {
    addUserStat(`Error fetching`)
    console.error('Error fetching /stats:', error);
  })

function addUserStat(text) {
  const span = document.createElement('span');
  span.className = 'mx-4 my-2';
  span.textContent = text;
  const userStats = document.getElementById('userStats');
  
  if (userStats) {
    userStats.appendChild(span);
  } else {
    console.error('Element with id "userStats" not found.');
  }
}

</script>

  <!-- Cloudflare Web Analytics --><script defer src='https://static.cloudflareinsights.com/beacon.min.js' data-cf-beacon='{"token": "4d7fcf12f8c140af87bcd51e98bdd4ae"}'></script><!-- End Cloudflare Web Analytics -->

</html>