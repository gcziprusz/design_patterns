class MyFile {
    protected name: string;
    constructor(name:string){
        this.name=name;
    }
    search (keyword :string){
        console.log(`\tSearching a file for keyword ${keyword} file name ${this.name}.\n`)
    }
    getName(){return this.name}
}


interface Component {
    search(keyworrd: string):void
}

class MyFolder {
    protected name: string;
    protected components: Component[]=[];
    constructor(name:string){
        this.name=name;
    }
    search (keyword :string){
        console.log(`Searching recursively for a keyword ${keyword} in folder name ${this.getName()}.\n`)
        for(let comp of this.components) {
            comp.search(keyword)
        }
    }
    getName(){return this.name}
    add(comp: Component){this.components.push(comp)}
}


const file1 = new MyFile("File1")
const file2 = new MyFile("File2")
const file3 = new MyFile( "File3")

const folder1 = new MyFolder("Folder1")

folder1.add(file1)

const folder2 = new MyFolder("Folder2")
folder2.add(file2)
folder2.add(file3)
folder2.add(folder1)

folder2.search("rose")

