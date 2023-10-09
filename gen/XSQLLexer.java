// Generated from /Users/koyeo/gobatis/gobatis/parser/xsql/XSQLLexer.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class XSQLLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMENT=1, EntityRef=2, SEA_WS=3, OPEN=4, TEXT=5, CLOSE=6, SPECIAL_CLOSE=7, 
		SLASH_CLOSE=8, SLASH=9, EQUALS=10, STRING=11, Name=12, S=13;
	public static final int
		INSIDE=1;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE", "INSIDE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMMENT", "EntityRef", "SEA_WS", "OPEN", "TEXT", "CLOSE", "SPECIAL_CLOSE", 
			"SLASH_CLOSE", "SLASH", "EQUALS", "STRING", "Name", "S", "HEXDIGIT", 
			"DIGIT", "NameChar", "NameStartChar"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'<'", null, "'>'", "'?>'", "'/>'", "'/'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMMENT", "EntityRef", "SEA_WS", "OPEN", "TEXT", "CLOSE", "SPECIAL_CLOSE", 
			"SLASH_CLOSE", "SLASH", "EQUALS", "STRING", "Name", "S"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public XSQLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "XSQLLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\r\u0087\u0006\uffff\uffff\u0006\uffff\uffff\u0002\u0000\u0007"+
		"\u0000\u0002\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007"+
		"\u0003\u0002\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007"+
		"\u0006\u0002\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n"+
		"\u0007\n\u0002\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002"+
		"\u000e\u0007\u000e\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0005"+
		"\u0000+\b\u0000\n\u0000\f\u0000.\t\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0003\u0002<\b\u0002\u0001"+
		"\u0002\u0004\u0002?\b\u0002\u000b\u0002\f\u0002@\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0004\u0004H\b\u0004\u000b\u0004"+
		"\f\u0004I\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\n\u0001\n\u0005\n`\b\n\n\n\f\nc\t\n\u0001\n\u0001\n\u0001\n\u0005\nh"+
		"\b\n\n\n\f\nk\t\n\u0001\n\u0003\nn\b\n\u0001\u000b\u0001\u000b\u0005\u000b"+
		"r\b\u000b\n\u000b\f\u000bu\t\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0083\b\u000f\u0001\u0010\u0003\u0010\u0086\b"+
		"\u0010\u0001,\u0000\u0011\u0002\u0001\u0004\u0002\u0006\u0003\b\u0004"+
		"\n\u0005\f\u0006\u000e\u0007\u0010\b\u0012\t\u0014\n\u0016\u000b\u0018"+
		"\f\u001a\r\u001c\u0000\u001e\u0000 \u0000\"\u0000\u0002\u0000\u0001\n"+
		"\u0002\u0000\t\t  \u0002\u0000&&<<\u0002\u0000\"\"<<\u0002\u0000\'\'<"+
		"<\u0003\u0000\t\n\r\r  \u0003\u000009AFaf\u0001\u000009\u0002\u0000-."+
		"__\u0003\u0000\u00b7\u00b7\u0300\u036f\u203f\u2040\b\u0000::AZaz\u2070"+
		"\u218f\u2c00\u2fef\u3001\u8000\ud7ff\u8000\uf900\u8000\ufdcf\u8000\ufdf0"+
		"\u8000\ufffd\u008d\u0000\u0002\u0001\u0000\u0000\u0000\u0000\u0004\u0001"+
		"\u0000\u0000\u0000\u0000\u0006\u0001\u0000\u0000\u0000\u0000\b\u0001\u0000"+
		"\u0000\u0000\u0000\n\u0001\u0000\u0000\u0000\u0001\f\u0001\u0000\u0000"+
		"\u0000\u0001\u000e\u0001\u0000\u0000\u0000\u0001\u0010\u0001\u0000\u0000"+
		"\u0000\u0001\u0012\u0001\u0000\u0000\u0000\u0001\u0014\u0001\u0000\u0000"+
		"\u0000\u0001\u0016\u0001\u0000\u0000\u0000\u0001\u0018\u0001\u0000\u0000"+
		"\u0000\u0001\u001a\u0001\u0000\u0000\u0000\u0002$\u0001\u0000\u0000\u0000"+
		"\u00045\u0001\u0000\u0000\u0000\u0006>\u0001\u0000\u0000\u0000\bB\u0001"+
		"\u0000\u0000\u0000\nG\u0001\u0000\u0000\u0000\fK\u0001\u0000\u0000\u0000"+
		"\u000eO\u0001\u0000\u0000\u0000\u0010T\u0001\u0000\u0000\u0000\u0012Y"+
		"\u0001\u0000\u0000\u0000\u0014[\u0001\u0000\u0000\u0000\u0016m\u0001\u0000"+
		"\u0000\u0000\u0018o\u0001\u0000\u0000\u0000\u001av\u0001\u0000\u0000\u0000"+
		"\u001cz\u0001\u0000\u0000\u0000\u001e|\u0001\u0000\u0000\u0000 \u0082"+
		"\u0001\u0000\u0000\u0000\"\u0085\u0001\u0000\u0000\u0000$%\u0005<\u0000"+
		"\u0000%&\u0005!\u0000\u0000&\'\u0005-\u0000\u0000\'(\u0005-\u0000\u0000"+
		"(,\u0001\u0000\u0000\u0000)+\t\u0000\u0000\u0000*)\u0001\u0000\u0000\u0000"+
		"+.\u0001\u0000\u0000\u0000,-\u0001\u0000\u0000\u0000,*\u0001\u0000\u0000"+
		"\u0000-/\u0001\u0000\u0000\u0000.,\u0001\u0000\u0000\u0000/0\u0005-\u0000"+
		"\u000001\u0005-\u0000\u000012\u0005>\u0000\u000023\u0001\u0000\u0000\u0000"+
		"34\u0006\u0000\u0000\u00004\u0003\u0001\u0000\u0000\u000056\u0005&\u0000"+
		"\u000067\u0003\u0018\u000b\u000078\u0005;\u0000\u00008\u0005\u0001\u0000"+
		"\u0000\u00009?\u0007\u0000\u0000\u0000:<\u0005\r\u0000\u0000;:\u0001\u0000"+
		"\u0000\u0000;<\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000=?\u0005"+
		"\n\u0000\u0000>9\u0001\u0000\u0000\u0000>;\u0001\u0000\u0000\u0000?@\u0001"+
		"\u0000\u0000\u0000@>\u0001\u0000\u0000\u0000@A\u0001\u0000\u0000\u0000"+
		"A\u0007\u0001\u0000\u0000\u0000BC\u0005<\u0000\u0000CD\u0001\u0000\u0000"+
		"\u0000DE\u0006\u0003\u0001\u0000E\t\u0001\u0000\u0000\u0000FH\b\u0001"+
		"\u0000\u0000GF\u0001\u0000\u0000\u0000HI\u0001\u0000\u0000\u0000IG\u0001"+
		"\u0000\u0000\u0000IJ\u0001\u0000\u0000\u0000J\u000b\u0001\u0000\u0000"+
		"\u0000KL\u0005>\u0000\u0000LM\u0001\u0000\u0000\u0000MN\u0006\u0005\u0002"+
		"\u0000N\r\u0001\u0000\u0000\u0000OP\u0005?\u0000\u0000PQ\u0005>\u0000"+
		"\u0000QR\u0001\u0000\u0000\u0000RS\u0006\u0006\u0002\u0000S\u000f\u0001"+
		"\u0000\u0000\u0000TU\u0005/\u0000\u0000UV\u0005>\u0000\u0000VW\u0001\u0000"+
		"\u0000\u0000WX\u0006\u0007\u0002\u0000X\u0011\u0001\u0000\u0000\u0000"+
		"YZ\u0005/\u0000\u0000Z\u0013\u0001\u0000\u0000\u0000[\\\u0005=\u0000\u0000"+
		"\\\u0015\u0001\u0000\u0000\u0000]a\u0005\"\u0000\u0000^`\b\u0002\u0000"+
		"\u0000_^\u0001\u0000\u0000\u0000`c\u0001\u0000\u0000\u0000a_\u0001\u0000"+
		"\u0000\u0000ab\u0001\u0000\u0000\u0000bd\u0001\u0000\u0000\u0000ca\u0001"+
		"\u0000\u0000\u0000dn\u0005\"\u0000\u0000ei\u0005\'\u0000\u0000fh\b\u0003"+
		"\u0000\u0000gf\u0001\u0000\u0000\u0000hk\u0001\u0000\u0000\u0000ig\u0001"+
		"\u0000\u0000\u0000ij\u0001\u0000\u0000\u0000jl\u0001\u0000\u0000\u0000"+
		"ki\u0001\u0000\u0000\u0000ln\u0005\'\u0000\u0000m]\u0001\u0000\u0000\u0000"+
		"me\u0001\u0000\u0000\u0000n\u0017\u0001\u0000\u0000\u0000os\u0003\"\u0010"+
		"\u0000pr\u0003 \u000f\u0000qp\u0001\u0000\u0000\u0000ru\u0001\u0000\u0000"+
		"\u0000sq\u0001\u0000\u0000\u0000st\u0001\u0000\u0000\u0000t\u0019\u0001"+
		"\u0000\u0000\u0000us\u0001\u0000\u0000\u0000vw\u0007\u0004\u0000\u0000"+
		"wx\u0001\u0000\u0000\u0000xy\u0006\f\u0003\u0000y\u001b\u0001\u0000\u0000"+
		"\u0000z{\u0007\u0005\u0000\u0000{\u001d\u0001\u0000\u0000\u0000|}\u0007"+
		"\u0006\u0000\u0000}\u001f\u0001\u0000\u0000\u0000~\u0083\u0003\"\u0010"+
		"\u0000\u007f\u0083\u0007\u0007\u0000\u0000\u0080\u0083\u0003\u001e\u000e"+
		"\u0000\u0081\u0083\u0007\b\u0000\u0000\u0082~\u0001\u0000\u0000\u0000"+
		"\u0082\u007f\u0001\u0000\u0000\u0000\u0082\u0080\u0001\u0000\u0000\u0000"+
		"\u0082\u0081\u0001\u0000\u0000\u0000\u0083!\u0001\u0000\u0000\u0000\u0084"+
		"\u0086\u0007\t\u0000\u0000\u0085\u0084\u0001\u0000\u0000\u0000\u0086#"+
		"\u0001\u0000\u0000\u0000\r\u0000\u0001,;>@Iaims\u0082\u0085\u0004\u0006"+
		"\u0000\u0000\u0005\u0001\u0000\u0004\u0000\u0000\u0000\u0001\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}