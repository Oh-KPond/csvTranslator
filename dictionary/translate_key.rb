require 'csv'
require 'json'

table = CSV.parse(File.read("dictionary/translation_key02192020.csv"), headers: true)

=begin
The structure I want to end up with in Go

map[string]map[string]string {
    "German": map[string]string {
        "germanWord":"englishWord",
        "germanWord2":"englishWord2",
    },
    "French": map[string]string {
        "frenchWord":"englishWord",
        "frenchWord2":"englishWord2",
    },
}
=end

# dataDict = { :nested_hash => { :first_key => 'Hello' } }
dataDict = {}

# make dataDict Language keys
table.headers.each do |header|
    dataDict[header] = {}
end

# make each Language key have map of {languageWord : englishWord}
dataDict.each do |language, innerMap|
    len = table.by_col["English"].length
    i = 0
    while i < len
        newKey = table[i][language]
        dataDict[language][newKey] = table[i]["English"]
        i += 1
    end
end

print(dataDict["German"])

=begin
After doing some research it seemed like the way to go 
was to turn this Ruby struct into json to then be used 
by the Go code.
=end

# write json to config file
File.open("dictionary/languageDict.json","w") do |f|
    f.write(JSON.pretty_generate(dataDict))
end