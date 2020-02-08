# -*- coding: utf-8 -*-
import spacy
import nltk
import gensim
import pickle
import csv

from nltk.corpus import wordnet as wn
from nltk.stem.wordnet import WordNetLemmatizer
from spacy.lang.en import English
from gensim import corpora


class TopicFinder:

    def __init__(self):
        spacy.load('en')
        self.parser = English()
        self.en_stop = set(nltk.corpus.stopwords.words('english'))

        nltk.download('wordnet')
        nltk.download('stopwords')

        self.ldamodel = gensim.models.ldamodel.LdaModel.load('ldaModel.gensim')

        self.dictionary = gensim.models.ldamodel.LdaModel.load('dictionary.gensim')

    # This tokenizes the input into a list of words
    def tokenize(self, text):
        lda_tokens = []
        tokens = self.parser(text)
        for token in tokens:
            if token.orth_.isspace():
                continue
            elif token.like_url:
                lda_tokens.append('URL')
            elif token.orth_.startswith('@'):
                lda_tokens.append('SCREEN_NAME')
            else:
                lda_tokens.append(token.lower_)
        return lda_tokens

    def get_lemma(self, word):
        lemma = wn.morphy(word)
        if lemma is None:
            return word
        else:
            return lemma

    def get_lemma2(self, word):
        return WordNetLemmatizer().lemmatize(word)

    def prepare_text_for_lda(self, text):
        tokens = self.tokenize(text)
        tokens = [token for token in tokens if len(token) > 4]
        tokens = [token for token in tokens if token not in self.en_stop]
        tokens = [self.get_lemma(token) for token in tokens]
        return tokens

    def toCSV(filename, csvFilename):
        with open(filename, 'r') as file:
            data = file.read().replace('\n', '')
            data = data.split()
            print(data)
            with open(csvFilename, 'w') as csvFile:
                writer = csv.writer(csvFile, delimiter=',')
                writer.writerow(data)
            csvFile.close()
        file.close()

    def train(self, filename, updateCurrentData=True):
        originalDictionaryData = []
        if updateCurrentData:
            originalDictionaryData = self.dictionary

        text_data = []
        with open(filename) as f:
            for line in f:
                tokens = self.prepare_text_for_lda(line)
                print(tokens)
                text_data.append(tokens)

        self.dictionary = corpora.Dictionary(text_data)
        self.dictionary.merge_with(originalDictionaryData)
        corpus = [self.dictionary.doc2bow(text) for text in text_data]
        print("\n this is the dictionary")
        print([self.dictionary.doc2bow(text) for text in text_data])
        print()
        pickle.dump(corpus, open('corpus.pkl', 'wb'))
        self.dictionary.save('dictionary.gensim')

        print()
        print("this is the gensim dictionary")
        print(self.dictionary)
        print()

        NUM_TOPICS = 1
        self.ldamodel = gensim.models.ldamodel.LdaModel(corpus, num_topics=NUM_TOPICS, id2word=self.dictionary,
                                                        passes=15)
        self.ldamodel.save('ldaModel.gensim')
        topics = self.ldamodel.print_topics(num_words=4)
        for topic in topics:
            print(topic)

    def getRelevancy(self, string, convID):
        string = self.prepare_text_for_lda(string)
        string_bow = self.dictionary.doc2bow(string)
        weightings = self.ldamodel.get_document_topics(string_bow)
        maxWeightAndID = weightings[0]
        for (converID, weight) in weightings:
            if weight > maxWeightAndID[1]:
                maxWeightAndID = (converID, weight)

        print("FIRST ONE: ", converID)
        print("CONVID: ", convID)
        return converID == convID

    def readPickle(self):
        pickleFile = open('corpus.pkl', "rb")  # statistics file is loaded
        self.statsData = pickle.load(pickleFile)
        print("hi")
        print(self.statsData)


topicFinder = TopicFinder()

# topicFinder.train('dataBrexitBBCData.csv')

new_doc = 'What advantages are there for the United Kingdom staying in the European Union'
print("FINAL ANSWER")
print(topicFinder.getRelevancy(new_doc, 0))
# topicFinder.readPickle()




