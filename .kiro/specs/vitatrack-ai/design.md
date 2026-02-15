# Design Document: VitaTrack AI

## Overview

VitaTrack AI is a full-stack web application that combines a Go backend with an Angular frontend to provide users with an intelligent health journaling system. The architecture follows a three-tier design pattern with clear separation between presentation (Angular), business logic (Go services), and data persistence (PostgreSQL + file storage).

The system enables users to upload health documents, track vital signs, and receive AI-powered insights through LLM integration. The calendar-based interface provides intuitive navigation through historical health data, while the AI analysis engine processes documents and generates personalized food recommendations.

### Key Design Principles

1. **Security First**: All health data is encrypted at rest and in transit, with strict authentication and authorization controls
2. **Scalability**: Stateless backend services enable horizontal scaling; file storage supports cloud backends
3. **Modularity**: Clear service boundaries allow independent development and testing of components
4. **Resilience**: Graceful degradation when LLM services are unavailable; retry logic for transient failures
5. **User Experience**: Optimistic UI updates and caching minimize perceived latency

## Architecture

### System Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                     Angular Frontend                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │   Calendar   │  │    Upload    │  │  Vitals      │      │
│  │   Component  │  │   Component  │  │  Component   │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│         │                  │                  │              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │           State Management (NgRx)                    │   │
│  └──────────────────────────────────────────────────────┘   │
│         │                  │                  │              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │           HTTP Client Service Layer                  │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                            │
                         HTTPS
                            │
┌─────────────────────────────────────────────────────────────┐
│                      Go Backend API                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │     Auth     │  │    Upload    │  │    Vitals    │      │
│  │   Handler    │  │   Handler    │  │   Handler    │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
│         │                  │                  │              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Service Layer                           │   │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐            │   │
│  │  │   Auth   │ │   File   │ │   LLM    │            │   │
│  │  │  Service │ │  Service │ │  Service │            │   │
│  │  └──────────┘ └──────────┘ └──────────┘            │   │
│  └──────────────────────────────────────────────────────┘   │
│         │                  │                  │              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │           Repository Layer                           │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
         │                  │                  │
         ▼                  ▼                  ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  PostgreSQL  │  │ File Storage │  │  LLM API     │
│   Database   │  │  (S3/Local)  │  │  (External)  │
└──────────────┘  └──────────────┘  └──────────────┘
```

### Technology Stack

**Frontend:**
- Angular 17+ with TypeScript
- NgRx for state management
- Angular Material for UI components
- Chart.js for data visualization
- RxJS for reactive programming

**Backend:**
- Go 1.21+ with standard library
- Gorilla Mux for HTTP routing
- GORM for database ORM
- JWT for authentication
- Multipart form handling for file uploads

**Data Storage:**
- PostgreSQL 15+ for structured data
- S3-compatible storage (AWS S3, MinIO) for files
- Redis for caching (optional)

**AI/ML:**
- OpenAI API or compatible LLM service
- Text extraction libraries (pdfcpu, tesseract-ocr)


## Components and Interfaces

### Backend Components

#### 1. HTTP Handlers

**AuthHandler**
```go
type AuthHandler struct {
    authService *AuthService
}

// Register creates a new user account
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request)

// Login authenticates a user and returns a JWT token
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request)

// Logout invalidates the user's session
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request)

// UpdateProfile updates user profile information
func (h *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request)

// DeleteAccount removes user account and all associated data
func (h *AuthHandler) DeleteAccount(w http.ResponseWriter, r *http.Request)
```

**FileHandler**
```go
type FileHandler struct {
    fileService *FileService
    llmService  *LLMService
}

// UploadFile handles multipart file upload
func (h *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request)

// GetFile retrieves a file by ID
func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request)

// DeleteFile removes a file
func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request)

// GetFilesByDate retrieves all files for a specific date
func (h *FileHandler) GetFilesByDate(w http.ResponseWriter, r *http.Request)

// GetCalendarMonth retrieves calendar entries for a month
func (h *FileHandler) GetCalendarMonth(w http.ResponseWriter, r *http.Request)
```

**VitalsHandler**
```go
type VitalsHandler struct {
    vitalsService *VitalsService
}

// CreateVital creates a new vital sign entry
func (h *VitalsHandler) CreateVital(w http.ResponseWriter, r *http.Request)

// GetVitals retrieves vitals for a date range
func (h *VitalsHandler) GetVitals(w http.ResponseWriter, r *http.Request)

// UpdateVital updates an existing vital sign entry
func (h *VitalsHandler) UpdateVital(w http.ResponseWriter, r *http.Request)

// DeleteVital removes a vital sign entry
func (h *VitalsHandler) DeleteVital(w http.ResponseWriter, r *http.Request)

// GetVitalsByType retrieves vitals filtered by type
func (h *VitalsHandler) GetVitalsByType(w http.ResponseWriter, r *http.Request)
```

**AnalysisHandler**
```go
type AnalysisHandler struct {
    llmService *LLMService
}

// GetAnalysis retrieves AI analysis for a date
func (h *AnalysisHandler) GetAnalysis(w http.ResponseWriter, r *http.Request)

// GetFoodRecommendations retrieves food recommendations
func (h *AnalysisHandler) GetFoodRecommendations(w http.ResponseWriter, r *http.Request)

// RefreshAnalysis triggers re-analysis of health data
func (h *AnalysisHandler) RefreshAnalysis(w http.ResponseWriter, r *http.Request)
```

#### 2. Service Layer

**AuthService**
```go
type AuthService struct {
    userRepo *UserRepository
    jwtSecret string
}

// CreateUser creates a new user with hashed password
func (s *AuthService) CreateUser(email, password, name string) (*User, error)

// AuthenticateUser validates credentials and returns JWT token
func (s *AuthService) AuthenticateUser(email, password string) (string, error)

// ValidateToken validates JWT token and returns user ID
func (s *AuthService) ValidateToken(token string) (string, error)

// UpdateUser updates user profile
func (s *AuthService) UpdateUser(userID string, updates map[string]interface{}) error

// DeleteUser removes user and all associated data
func (s *AuthService) DeleteUser(userID string) error
```

**FileService**
```go
type FileService struct {
    fileRepo    *FileRepository
    storage     StorageBackend
    textExtractor TextExtractor
}

// StoreFile saves file to storage and creates database record
func (s *FileService) StoreFile(userID string, file multipart.File, filename string, date time.Time) (*HealthDocument, error)

// GetFile retrieves file metadata and content
func (s *FileService) GetFile(userID, fileID string) (*HealthDocument, []byte, error)

// DeleteFile removes file from storage and database
func (s *FileService) DeleteFile(userID, fileID string) error

// GetFilesByDate retrieves all files for a specific date
func (s *FileService) GetFilesByDate(userID string, date time.Time) ([]*HealthDocument, error)

// GetCalendarEntries retrieves calendar entries for a month
func (s *FileService) GetCalendarEntries(userID string, year, month int) (map[string]*CalendarEntry, error)

// ExtractText extracts text content from a file
func (s *FileService) ExtractText(file []byte, mimeType string) (string, error)
```

**VitalsService**
```go
type VitalsService struct {
    vitalsRepo *VitalsRepository
}

// CreateVital creates a new vital sign entry
func (s *VitalsService) CreateVital(userID string, vitalType string, value float64, date time.Time) (*VitalSign, error)

// GetVitals retrieves vitals for a date range
func (s *VitalsService) GetVitals(userID string, startDate, endDate time.Time) ([]*VitalSign, error)

// UpdateVital updates an existing vital sign
func (s *VitalsService) UpdateVital(userID, vitalID string, value float64) error

// DeleteVital removes a vital sign entry
func (s *VitalsService) DeleteVital(userID, vitalID string) error

// ValidateVital validates vital sign value based on type
func (s *VitalsService) ValidateVital(vitalType string, value float64) error
```

**LLMService**
```go
type LLMService struct {
    client      LLMClient
    cache       Cache
    queue       *AnalysisQueue
    vitalsRepo  *VitalsRepository
    fileRepo    *FileRepository
}

// AnalyzeHealthDocument sends document text to LLM for analysis
func (s *LLMService) AnalyzeHealthDocument(userID string, documentText string, date time.Time) (*HealthInsight, error)

// GenerateFoodRecommendations generates dietary recommendations
func (s *LLMService) GenerateFoodRecommendations(userID string, date time.Time) ([]*FoodRecommendation, error)

// QueueAnalysis adds analysis request to queue
func (s *LLMService) QueueAnalysis(userID, documentID string) error

// ProcessQueue processes queued analysis requests
func (s *LLMService) ProcessQueue() error

// BuildHealthContext gathers relevant health data for LLM prompt
func (s *LLMService) BuildHealthContext(userID string, date time.Time) (string, error)
```

#### 3. Repository Layer

**UserRepository**
```go
type UserRepository struct {
    db *gorm.DB
}

func (r *UserRepository) Create(user *User) error
func (r *UserRepository) FindByEmail(email string) (*User, error)
func (r *UserRepository) FindByID(id string) (*User, error)
func (r *UserRepository) Update(user *User) error
func (r *UserRepository) Delete(id string) error
```

**FileRepository**
```go
type FileRepository struct {
    db *gorm.DB
}

func (r *FileRepository) Create(doc *HealthDocument) error
func (r *FileRepository) FindByID(id string) (*HealthDocument, error)
func (r *FileRepository) FindByUserAndDate(userID string, date time.Time) ([]*HealthDocument, error)
func (r *FileRepository) FindByUserAndDateRange(userID string, start, end time.Time) ([]*HealthDocument, error)
func (r *FileRepository) Delete(id string) error
func (r *FileRepository) GetCalendarEntries(userID string, year, month int) (map[string]*CalendarEntry, error)
```

**VitalsRepository**
```go
type VitalsRepository struct {
    db *gorm.DB
}

func (r *VitalsRepository) Create(vital *VitalSign) error
func (r *VitalsRepository) FindByID(id string) (*VitalSign, error)
func (r *VitalsRepository) FindByUserAndDateRange(userID string, start, end time.Time) ([]*VitalSign, error)
func (r *VitalsRepository) FindByUserAndType(userID, vitalType string, start, end time.Time) ([]*VitalSign, error)
func (r *VitalsRepository) Update(vital *VitalSign) error
func (r *VitalsRepository) Delete(id string) error
```

**AnalysisRepository**
```go
type AnalysisRepository struct {
    db *gorm.DB
}

func (r *AnalysisRepository) CreateInsight(insight *HealthInsight) error
func (r *AnalysisRepository) CreateRecommendation(rec *FoodRecommendation) error
func (r *AnalysisRepository) FindInsightsByDate(userID string, date time.Time) ([]*HealthInsight, error)
func (r *AnalysisRepository) FindRecommendationsByDate(userID string, date time.Time) ([]*FoodRecommendation, error)
func (r *AnalysisRepository) DeleteInsightsByDocument(documentID string) error
```

### Frontend Components

#### Angular Components

**CalendarComponent**
```typescript
@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html'
})
export class CalendarComponent {
  currentMonth: Date;
  calendarDays: CalendarDay[];
  selectedDate: Date;
  
  ngOnInit(): void;
  navigateMonth(direction: number): void;
  selectDate(date: Date): void;
  hasEntries(date: Date): boolean;
}
```

**DateDetailComponent**
```typescript
@Component({
  selector: 'app-date-detail',
  templateUrl: './date-detail.component.html'
})
export class DateDetailComponent {
  selectedDate: Date;
  documents: HealthDocument[];
  vitals: VitalSign[];
  insights: HealthInsight[];
  recommendations: FoodRecommendation[];
  
  ngOnInit(): void;
  loadDateData(): void;
  deleteDocument(id: string): void;
  refreshAnalysis(): void;
}
```

**UploadComponent**
```typescript
@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html'
})
export class UploadComponent {
  selectedFile: File;
  targetDate: Date;
  uploadProgress: number;
  
  onFileSelected(event: Event): void;
  uploadFile(): void;
  validateFile(file: File): boolean;
}
```

**VitalsFormComponent**
```typescript
@Component({
  selector: 'app-vitals-form',
  templateUrl: './vitals-form.component.html'
})
export class VitalsFormComponent {
  vitalTypes: string[];
  selectedType: string;
  value: number;
  date: Date;
  
  onSubmit(): void;
  validateValue(): boolean;
}
```

**VitalsChartComponent**
```typescript
@Component({
  selector: 'app-vitals-chart',
  templateUrl: './vitals-chart.component.html'
})
export class VitalsChartComponent {
  chartData: ChartData;
  selectedVitalType: string;
  dateRange: DateRange;
  
  ngOnInit(): void;
  loadChartData(): void;
  updateChart(): void;
  exportData(): void;
}
```

#### Angular Services

**ApiService**
```typescript
@Injectable({ providedIn: 'root' })
export class ApiService {
  constructor(private http: HttpClient) {}
  
  // Auth endpoints
  register(email: string, password: string, name: string): Observable<User>;
  login(email: string, password: string): Observable<AuthResponse>;
  logout(): Observable<void>;
  
  // File endpoints
  uploadFile(file: File, date: Date): Observable<HealthDocument>;
  getFile(id: string): Observable<Blob>;
  deleteFile(id: string): Observable<void>;
  getFilesByDate(date: Date): Observable<HealthDocument[]>;
  getCalendarMonth(year: number, month: number): Observable<CalendarEntry[]>;
  
  // Vitals endpoints
  createVital(vital: VitalSignInput): Observable<VitalSign>;
  getVitals(startDate: Date, endDate: Date): Observable<VitalSign[]>;
  updateVital(id: string, value: number): Observable<VitalSign>;
  deleteVital(id: string): Observable<void>;
  
  // Analysis endpoints
  getAnalysis(date: Date): Observable<HealthInsight[]>;
  getFoodRecommendations(date: Date): Observable<FoodRecommendation[]>;
  refreshAnalysis(date: Date): Observable<void>;
}
```

**StateService (NgRx)**
```typescript
// Actions
export const loadCalendarMonth = createAction('[Calendar] Load Month', props<{ year: number, month: number }>());
export const uploadFile = createAction('[Upload] Upload File', props<{ file: File, date: Date }>());
export const createVital = createAction('[Vitals] Create Vital', props<{ vital: VitalSignInput }>());

// Selectors
export const selectCalendarEntries = createSelector(selectCalendarState, state => state.entries);
export const selectSelectedDate = createSelector(selectCalendarState, state => state.selectedDate);
export const selectVitalsData = createSelector(selectVitalsState, state => state.vitals);

// Effects
@Injectable()
export class CalendarEffects {
  loadMonth$ = createEffect(() => this.actions$.pipe(
    ofType(loadCalendarMonth),
    switchMap(action => this.api.getCalendarMonth(action.year, action.month))
  ));
}
```


## Data Models

### Database Schema

**Users Table**
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
```

**Health Documents Table**
```sql
CREATE TABLE health_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    filename VARCHAR(255) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    file_size BIGINT NOT NULL,
    storage_path VARCHAR(500) NOT NULL,
    upload_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_health_docs_user_date ON health_documents(user_id, upload_date);
CREATE INDEX idx_health_docs_user ON health_documents(user_id);
```

**Vital Signs Table**
```sql
CREATE TABLE vital_signs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vital_type VARCHAR(50) NOT NULL,
    value DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20) NOT NULL,
    measurement_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_vitals_user_date ON vital_signs(user_id, measurement_date);
CREATE INDEX idx_vitals_user_type ON vital_signs(user_id, vital_type);
```

**Health Insights Table**
```sql
CREATE TABLE health_insights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    document_id UUID REFERENCES health_documents(id) ON DELETE CASCADE,
    insight_text TEXT NOT NULL,
    insight_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES health_documents(id)
);

CREATE INDEX idx_insights_user_date ON health_insights(user_id, insight_date);
CREATE INDEX idx_insights_document ON health_insights(document_id);
```

**Food Recommendations Table**
```sql
CREATE TABLE food_recommendations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recommendation_text TEXT NOT NULL,
    reasoning TEXT NOT NULL,
    recommendation_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_recommendations_user_date ON food_recommendations(user_id, recommendation_date);
```

### Go Data Models

```go
type User struct {
    ID           string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Email        string    `gorm:"uniqueIndex;not null" json:"email"`
    PasswordHash string    `gorm:"not null" json:"-"`
    Name         string    `gorm:"not null" json:"name"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type HealthDocument struct {
    ID               string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    UserID           string    `gorm:"type:uuid;not null" json:"user_id"`
    Filename         string    `gorm:"not null" json:"filename"`
    OriginalFilename string    `gorm:"not null" json:"original_filename"`
    MimeType         string    `gorm:"not null" json:"mime_type"`
    FileSize         int64     `gorm:"not null" json:"file_size"`
    StoragePath      string    `gorm:"not null" json:"-"`
    UploadDate       time.Time `gorm:"type:date;not null" json:"upload_date"`
    CreatedAt        time.Time `json:"created_at"`
}

type VitalSign struct {
    ID              string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    UserID          string    `gorm:"type:uuid;not null" json:"user_id"`
    VitalType       string    `gorm:"not null" json:"vital_type"`
    Value           float64   `gorm:"type:decimal(10,2);not null" json:"value"`
    Unit            string    `gorm:"not null" json:"unit"`
    MeasurementDate time.Time `gorm:"type:date;not null" json:"measurement_date"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type HealthInsight struct {
    ID          string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
    DocumentID  *string   `gorm:"type:uuid" json:"document_id,omitempty"`
    InsightText string    `gorm:"type:text;not null" json:"insight_text"`
    InsightDate time.Time `gorm:"type:date;not null" json:"insight_date"`
    CreatedAt   time.Time `json:"created_at"`
}

type FoodRecommendation struct {
    ID                 string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    UserID             string    `gorm:"type:uuid;not null" json:"user_id"`
    RecommendationText string    `gorm:"type:text;not null" json:"recommendation_text"`
    Reasoning          string    `gorm:"type:text;not null" json:"reasoning"`
    RecommendationDate time.Time `gorm:"type:date;not null" json:"recommendation_date"`
    CreatedAt          time.Time `json:"created_at"`
}

type CalendarEntry struct {
    Date            time.Time             `json:"date"`
    Documents       []*HealthDocument     `json:"documents"`
    Vitals          []*VitalSign          `json:"vitals"`
    Insights        []*HealthInsight      `json:"insights"`
    Recommendations []*FoodRecommendation `json:"recommendations"`
}
```

### TypeScript Data Models

```typescript
export interface User {
  id: string;
  email: string;
  name: string;
  created_at: Date;
  updated_at: Date;
}

export interface HealthDocument {
  id: string;
  user_id: string;
  filename: string;
  original_filename: string;
  mime_type: string;
  file_size: number;
  upload_date: Date;
  created_at: Date;
}

export interface VitalSign {
  id: string;
  user_id: string;
  vital_type: VitalType;
  value: number;
  unit: string;
  measurement_date: Date;
  created_at: Date;
  updated_at: Date;
}

export enum VitalType {
  BLOOD_PRESSURE = 'blood_pressure',
  HEART_RATE = 'heart_rate',
  WEIGHT = 'weight',
  TEMPERATURE = 'temperature',
  BLOOD_GLUCOSE = 'blood_glucose',
  OXYGEN_SATURATION = 'oxygen_saturation'
}

export interface HealthInsight {
  id: string;
  user_id: string;
  document_id?: string;
  insight_text: string;
  insight_date: Date;
  created_at: Date;
}

export interface FoodRecommendation {
  id: string;
  user_id: string;
  recommendation_text: string;
  reasoning: string;
  recommendation_date: Date;
  created_at: Date;
}

export interface CalendarEntry {
  date: Date;
  documents: HealthDocument[];
  vitals: VitalSign[];
  insights: HealthInsight[];
  recommendations: FoodRecommendation[];
}

export interface CalendarDay {
  date: Date;
  hasEntries: boolean;
  isCurrentMonth: boolean;
  isToday: boolean;
}
```

## API Design

### REST API Endpoints

**Authentication Endpoints**

```
POST /api/v1/auth/register
Request: { "email": string, "password": string, "name": string }
Response: { "user": User, "token": string }

POST /api/v1/auth/login
Request: { "email": string, "password": string }
Response: { "token": string, "user": User }

POST /api/v1/auth/logout
Headers: Authorization: Bearer <token>
Response: { "message": "Logged out successfully" }

PUT /api/v1/auth/profile
Headers: Authorization: Bearer <token>
Request: { "name": string }
Response: { "user": User }

DELETE /api/v1/auth/account
Headers: Authorization: Bearer <token>
Response: { "message": "Account deleted successfully" }
```

**File Management Endpoints**

```
POST /api/v1/files/upload
Headers: Authorization: Bearer <token>
Content-Type: multipart/form-data
Request: { "file": File, "date": string (YYYY-MM-DD) }
Response: { "document": HealthDocument }

GET /api/v1/files/:id
Headers: Authorization: Bearer <token>
Response: File binary data with appropriate Content-Type

DELETE /api/v1/files/:id
Headers: Authorization: Bearer <token>
Response: { "message": "File deleted successfully" }

GET /api/v1/files/date/:date
Headers: Authorization: Bearer <token>
Response: { "documents": HealthDocument[] }

GET /api/v1/calendar/:year/:month
Headers: Authorization: Bearer <token>
Response: { "entries": { [date: string]: CalendarEntry } }
```

**Vitals Endpoints**

```
POST /api/v1/vitals
Headers: Authorization: Bearer <token>
Request: { "vital_type": string, "value": number, "date": string }
Response: { "vital": VitalSign }

GET /api/v1/vitals?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
Headers: Authorization: Bearer <token>
Response: { "vitals": VitalSign[] }

GET /api/v1/vitals/type/:type?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
Headers: Authorization: Bearer <token>
Response: { "vitals": VitalSign[] }

PUT /api/v1/vitals/:id
Headers: Authorization: Bearer <token>
Request: { "value": number }
Response: { "vital": VitalSign }

DELETE /api/v1/vitals/:id
Headers: Authorization: Bearer <token>
Response: { "message": "Vital deleted successfully" }
```

**Analysis Endpoints**

```
GET /api/v1/analysis/:date
Headers: Authorization: Bearer <token>
Response: { "insights": HealthInsight[] }

GET /api/v1/recommendations/:date
Headers: Authorization: Bearer <token>
Response: { "recommendations": FoodRecommendation[] }

POST /api/v1/analysis/refresh/:date
Headers: Authorization: Bearer <token>
Response: { "message": "Analysis queued successfully" }
```

### Error Response Format

All error responses follow this structure:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": {}
  }
}
```

Common HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request (validation errors)
- 401: Unauthorized (missing or invalid token)
- 403: Forbidden (insufficient permissions)
- 404: Not Found
- 413: Payload Too Large (file size exceeded)
- 500: Internal Server Error


## Correctness Properties

A property is a characteristic or behavior that should hold true across all valid executions of a system—essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.

### Authentication and User Management Properties

**Property 1: User registration creates valid profiles**
*For any* valid registration data (email, password, name), creating a user should result in a stored user profile with a unique ID and hashed password.
**Validates: Requirements 1.1**

**Property 2: Valid credentials authenticate successfully**
*For any* registered user, providing their correct email and password should return a valid JWT token that can be used for subsequent API calls.
**Validates: Requirements 1.2**

**Property 3: Invalid tokens are rejected**
*For any* malformed, expired, or non-existent token, API requests using that token should be rejected with a 401 Unauthorized response.
**Validates: Requirements 1.3**

**Property 4: Passwords are never stored in plaintext**
*For any* user registration or password update, the stored password hash should never equal the plaintext password.
**Validates: Requirements 1.4**

**Property 5: Profile updates persist correctly**
*For any* user and valid profile update, updating the profile then retrieving it should reflect the new values.
**Validates: Requirements 1.5**

**Property 6: Account deletion cascades to all user data**
*For any* user with health documents, vitals, and calendar entries, deleting the account should remove all associated data from the database and file storage.
**Validates: Requirements 1.6**

### File Upload and Storage Properties

**Property 7: Valid file uploads succeed and are stored**
*For any* valid file (supported format, under size limit) and target date, uploading the file should result in a stored document with a unique ID and correct metadata.
**Validates: Requirements 2.1**

**Property 8: Unsupported file formats are rejected**
*For any* file with an unsupported MIME type, the upload should be rejected with a format error.
**Validates: Requirements 2.3**

**Property 9: File upload creates calendar entry**
*For any* successful file upload to a date, the calendar entry for that date should contain the uploaded document.
**Validates: Requirements 2.4**

**Property 10: Multiple uploads accumulate in calendar entry**
*For any* date, uploading multiple files to that date should result in all files being present in the calendar entry.
**Validates: Requirements 2.5**

**Property 11: Document IDs are unique**
*For any* set of uploaded documents, all document IDs should be unique.
**Validates: Requirements 2.6**

**Property 12: File upload and download round trip**
*For any* uploaded file, downloading it should return data identical to the original upload.
**Validates: Requirements 4.2**

**Property 13: File deletion removes from storage and calendar**
*For any* uploaded document, deleting it should remove it from file storage and from the associated calendar entry.
**Validates: Requirements 4.3**

**Property 14: Empty calendar entries are removed**
*For any* calendar entry with a single document, deleting that document should remove the calendar entry from the database.
**Validates: Requirements 4.4**

**Property 15: Users can only access their own documents**
*For any* two different users, attempting to access another user's documents should return a 403 Forbidden response.
**Validates: Requirements 4.5**

**Property 16: File paths follow structured hierarchy**
*For any* uploaded document, the storage path should follow the pattern: {userID}/{date}/{uniqueFilename}.
**Validates: Requirements 12.1**

**Property 17: Filenames are unique and preserve extensions**
*For any* set of uploaded files, all generated filenames should be unique, and each should preserve the original file extension.
**Validates: Requirements 12.2**

### Calendar and Navigation Properties

**Property 18: Calendar displays entries for correct dates**
*For any* month and year, the calendar view should show visual indicators only for dates that have calendar entries.
**Validates: Requirements 3.1**

**Property 19: Date selection displays all documents**
*For any* date with calendar entries, selecting that date should display all documents uploaded to that date.
**Validates: Requirements 3.2**

**Property 20: Month navigation loads correct data**
*For any* month transition, navigating to a new month should load and display calendar entries only for that month.
**Validates: Requirements 3.4**

### Vital Signs Properties

**Property 21: Valid vitals are stored correctly**
*For any* valid vital sign (supported type, valid value, date), creating the vital should result in it being stored in the database with correct metadata.
**Validates: Requirements 5.1**

**Property 22: Invalid vital values are rejected**
*For any* vital sign with an invalid value (negative, out of range), the creation should be rejected with a validation error.
**Validates: Requirements 5.3**

**Property 23: Vital queries return correct date range**
*For any* date range query, the returned vitals should all have measurement dates within the specified range and be sorted by date.
**Validates: Requirements 5.4**

**Property 24: Vital updates persist correctly**
*For any* existing vital sign, updating its value then retrieving it should reflect the new value.
**Validates: Requirements 5.5**

**Property 25: Vital deletion removes from database**
*For any* vital sign, deleting it should result in it no longer being retrievable from the database.
**Validates: Requirements 5.6**

### AI Analysis and LLM Integration Properties

**Property 26: Document upload triggers LLM analysis**
*For any* uploaded health document, the system should extract text and send it to the LLM service for analysis.
**Validates: Requirements 6.1**

**Property 27: LLM insights are stored with calendar entry**
*For any* LLM analysis response, the generated insights should be stored in the database associated with the correct calendar entry and date.
**Validates: Requirements 6.2**

**Property 28: Calendar entry displays all insights**
*For any* calendar entry with insights, viewing the entry should display all associated health insights.
**Validates: Requirements 6.3**

**Property 29: LLM requests include relevant vitals**
*For any* LLM analysis request for a date, the request context should include all vital signs from that date.
**Validates: Requirements 6.6**

**Property 30: Food recommendations include reasoning**
*For any* generated food recommendation, it should include non-empty reasoning text explaining why it was suggested.
**Validates: Requirements 7.4**

**Property 31: Recommendation refresh uses latest data**
*For any* date with existing recommendations, adding new health data then refreshing should generate different recommendations that incorporate the new data.
**Validates: Requirements 7.6**

**Property 32: Failed LLM requests are queued**
*For any* LLM request that fails due to service unavailability, the request should be added to a queue for later processing.
**Validates: Requirements 10.3**

**Property 33: LLM retries use exponential backoff**
*For any* failed LLM request, retry attempts should occur with exponentially increasing delays.
**Validates: Requirements 10.4**

**Property 34: Large inputs are chunked**
*For any* LLM request exceeding the token limit, the input should be split into chunks, processed separately, and results combined.
**Validates: Requirements 13.3**

**Property 35: LLM responses are cached**
*For any* LLM request, making the same request twice should return the cached result without calling the LLM API again.
**Validates: Requirements 13.4**

**Property 36: LLM requests timeout appropriately**
*For any* LLM request that takes longer than 30 seconds, the request should timeout and return an error.
**Validates: Requirements 13.5**

**Property 37: LLM calls are logged**
*For any* LLM request, both the request and response should be logged for debugging purposes.
**Validates: Requirements 13.6**

**Property 38: LLM responses are validated**
*For any* LLM response, the system should validate and sanitize the content before storing it in the database.
**Validates: Requirements 13.7**

### Data Visualization Properties

**Property 39: Chart data matches vitals data**
*For any* vital type and date range, the chart data should exactly match the vitals stored in the database for that type and range.
**Validates: Requirements 8.1**

**Property 40: Date range changes update chart**
*For any* chart with a date range, changing the range should update the displayed data to match the new range.
**Validates: Requirements 8.5**

**Property 41: CSV export contains correct data**
*For any* chart, exporting to CSV should produce a file containing all the data points displayed in the chart.
**Validates: Requirements 8.6**

### Security and Privacy Properties

**Property 42: Files are encrypted at rest**
*For any* uploaded file, the stored file content should be encrypted and not equal to the plaintext original.
**Validates: Requirements 9.1**

**Property 43: Malicious inputs are sanitized**
*For any* user input containing SQL injection or XSS patterns, the input should be sanitized or rejected before processing.
**Validates: Requirements 9.3**

**Property 44: Rate limiting throttles excessive requests**
*For any* user making more than the allowed number of requests per time window, subsequent requests should be rejected with a 429 Too Many Requests response.
**Validates: Requirements 9.4**

**Property 45: Document access is logged**
*For any* document access operation (view, download, delete), an audit log entry should be created with user ID, document ID, and timestamp.
**Validates: Requirements 9.5**

**Property 46: Session expiry clears cached data**
*For any* expired session, the frontend should clear all cached health data and require re-authentication for subsequent requests.
**Validates: Requirements 9.6**

**Property 47: LLM requests minimize PII**
*For any* LLM request, personally identifiable information (email, full name, address) should be filtered out unless necessary for analysis.
**Validates: Requirements 9.7**

### API Design Properties

**Property 48: API responses follow standard format**
*For any* API response (success or error), the response should follow the standardized JSON format with consistent structure.
**Validates: Requirements 11.2**

**Property 49: Malformed requests return 400**
*For any* API request with invalid JSON or missing required fields, the response should be 400 Bad Request with validation details.
**Validates: Requirements 11.3**

**Property 50: Unauthenticated requests return 401**
*For any* protected API endpoint, requests without a valid authentication token should return 401 Unauthorized.
**Validates: Requirements 11.4**

**Property 51: Unauthorized access returns 403**
*For any* API request attempting to access resources owned by another user, the response should be 403 Forbidden.
**Validates: Requirements 11.5**

### Frontend State Management Properties

**Property 52: Upload updates calendar state**
*For any* file upload, the frontend calendar state should be updated to include the new document without requiring a page refresh.
**Validates: Requirements 14.1**

**Property 53: Vital entry updates visualization**
*For any* new vital sign entry, the frontend visualization state should immediately reflect the new data point.
**Validates: Requirements 14.2**

**Property 54: Calendar data is cached**
*For any* calendar entry, requesting the same entry multiple times should use cached data and not make redundant API calls.
**Validates: Requirements 14.3**

**Property 55: Failed actions rollback state**
*For any* optimistic UI update that fails, the frontend state should rollback to the previous state before the action.
**Validates: Requirements 14.5**

### Search and Filtering Properties

**Property 56: Search returns matching documents**
*For any* search query, all returned documents should contain the search term in either their metadata or associated AI analysis content.
**Validates: Requirements 15.1**

**Property 57: Date filters return correct range**
*For any* date range filter, all returned calendar entries should have dates within the specified range.
**Validates: Requirements 15.2**

**Property 58: Type filters return correct vitals**
*For any* vital type filter, all returned entries should contain at least one vital sign of the specified type.
**Validates: Requirements 15.3**

**Property 59: Search results are sorted correctly**
*For any* search query, results should be sorted first by relevance score (descending) then by date (descending).
**Validates: Requirements 15.4**

