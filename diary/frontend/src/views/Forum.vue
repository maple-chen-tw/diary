<template>
    <div class="forum-question-container">
      <div class="question-header">
        <h1>{{ question.title }}</h1>
        <p class="question-date">{{ formatDate(question.create_at) }}</p>
        <p>{{ question.content }}</p>
  
        <div class="question-actions">
          <button @click="voteQuestion('like')" class="btn btn-success" :disabled="hasVoted">Like ({{ question.likes }})</button>
          <button @click="voteQuestion('dislike')" class="btn btn-danger" :disabled="hasVoted">Dislike ({{ question.dislikes }})</button>
        </div>
      </div>
  
      <div class="answers">
        <h3>Answers</h3>
        <div v-for="answer in answers" :key="answer.id" class="answer-item">
          <p>{{ answer.content }}</p>
          <p class="answer-date">{{ formatDate(answer.create_at) }}</p>
        </div>
  
        <div class="answer-form">
          <textarea v-model="newAnswer" placeholder="Write your answer..."></textarea>
          <button @click="submitAnswer" class="btn btn-secondary">Post Answer</button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        question: {},       // The question data
        answers: [],        // Answers to the question
        newAnswer: '',      // New answer input
        hasVoted: false,    // Whether the user has voted
      };
    },
    methods: {
      async fetchQuestion() {
        const questionId = this.$route.params.id; // Get question ID from route params
        try {
          const response = await axios.get(`/public_questions/${questionId}`);
          this.question = response.data;
          this.checkUserVote(questionId);
          this.fetchAnswers(questionId);
        } catch (error) {
          console.error('Error fetching question:', error);
        }
      },
  
      async fetchAnswers(questionId) {
        try {
          const response = await axios.get(`/public_questions/${questionId}/answers`);
          this.answers = response.data;
        } catch (error) {
          console.error('Error fetching answers:', error);
        }
      },
  
      async submitAnswer() {
        if (!this.newAnswer.trim()) {
          alert('Please write an answer');
          return;
        }
  
        const questionId = this.$route.params.id;
        try {
          const response = await axios.post(`/public_questions/${questionId}/answers`, {
            content: this.newAnswer,
          });
          this.answers.push(response.data);  // Add new answer to the list
          this.newAnswer = '';  // Clear the input field
        } catch (error) {
          console.error('Error submitting answer:', error);
        }
      },
  
      async voteQuestion(voteType) {
        const questionId = this.$route.params.id;
        try {
          await axios.post(`/public_questions/${questionId}/vote`, {
            vote_type: voteType,
          });
          if (voteType === 'like') {
            this.question.likes += 1;
          } else if (voteType === 'dislike') {
            this.question.dislikes += 1;
          }
          this.hasVoted = true;
        } catch (error) {
          console.error('Error voting on question:', error);
        }
      },
  
      async checkUserVote(questionId) {
        try {
          const response = await axios.get(`/public_questions/${questionId}/vote/status`);
          this.hasVoted = response.data.hasVoted;
        } catch (error) {
          console.error('Error checking vote status:', error);
        }
      },
  
      formatDate(dateString) {
        const date = new Date(dateString);
        return date.toLocaleDateString();
      },
    },
  
    mounted() {
      this.fetchQuestion();
    },
  };
  </script>
  
  <style scoped>
  .forum-question-container {
    max-width: 900px;
    margin: 50px auto;
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  
  .question-header {
    margin-bottom: 30px;
  }
  
  .question-header h1 {
    font-size: 26px;
    color: #333;
  }
  
  .question-date {
    font-size: 14px;
    color: #999;
    margin-bottom: 10px;
  }
  
  .question-actions {
    display: flex;
    gap: 10px;
    margin-top: 10px;
  }
  
  .answers {
    margin-top: 30px;
  }
  
  .answer-item {
    background-color: #f1f1f1;
    padding: 15px;
    margin-bottom: 20px;
    border-radius: 8px;
  }
  
  .answer-date {
    font-size: 12px;
    color: #999;
  }
  
  .answer-form textarea {
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border-radius: 8px;
    border: 1px solid #ccc;
    margin-bottom: 10px;
  }
  
  .answer-form button {
    padding: 10px 20px;
    font-size: 16px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    background-color: #6c757d;
    color: white;
  }
  
  .answer-form button:hover {
    opacity: 0.9;
  }
  
  .btn-success {
    background-color: #28a745;
    color: white;
  }
  
  .btn-danger {
    background-color: #dc3545;
    color: white;
  }
  
  .btn-secondary {
    background-color: #6c757d;
    color: white;
  }
  
  button:hover {
    opacity: 0.9;
  }
  
  button:focus {
    outline: none;
  }
  </style>
  