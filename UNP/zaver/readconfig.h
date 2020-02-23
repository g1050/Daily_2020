#ifndef _READCONFIG_H_
#define _READCONFIG_H_
#include "debug.h"
#include <map>
#include <string>
#include <iostream>
#include <fstream>
#define COMMENT_CHAR '#'
using namespace std;

class ReadConfig{
public:
    ReadConfig(const string& filename)
        :m_filename(filename){}
public:
    bool readFile(){
        m_map.clear();
        ifstream in(m_filename,ios::in);
        if(!in) return false;
        string line,key,value;

        while(getline(in,line)){
            if(analyseLine(line,key,value)){
                m_map[key] = value;
            }else{
                return false;
            }
        }

        in.close();
        return true;
    }

    void print(){
        map<string,string>::const_iterator it = m_map.begin();
        for(;it != m_map.end();it++){
            cout << it->first << " " << it->second << endl;
        }
    }

private:
    bool analyseLine(const string& line,string& key,string &value){

        if(line.empty()) return false;

        size_t start_pos = 0,end_pos = line.size()-1,cur_pos = 0;
        //Check if there is '#' 
        if((cur_pos = line.find(COMMENT_CHAR)) == 0){
            //first char is '#',the whole line is comment
            return true;
        }else if(cur_pos == string::npos){
            ;
        }else{
            end_pos = cur_pos - 1;
        }
        //delete the string after first #
        string new_line = line.substr(start_pos,end_pos - start_pos + 1);
        if((cur_pos = new_line.find('=')) == string::npos)
            return false;
        key = new_line.substr(start_pos,cur_pos);
        value = new_line.substr(cur_pos + 1,end_pos - cur_pos);

        //trim the space
        trim(key);
        if(key.empty())
            return false;
        trim(value);
        return true;
    }

    void trim(string &s){
        if(s.empty())
            return ;
        size_t i ;
        //before content
        for(i = 0;i<s.size();i++){
            if(s[i] != ' ' && s[i] != '\t')
                break;
        }
        s = s.substr(i);
        //after content
        for(i = 0;i<s.size();i++){
            if(i == ' '|| s[i] == '\t')
                break;
        }
        s = s.substr(0,i);
    }

private:
    map<string,string> m_map;
    string m_filename;
};
#endif
